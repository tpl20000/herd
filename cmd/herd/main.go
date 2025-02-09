package main

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"time"

	"github.com/seveas/herd"
	"github.com/seveas/herd/scripting"

	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var currentUser *userData

type userData struct {
	user            *user.User
	cacheDir        string
	configDir       string
	systemConfigDir string
	dataDir         string
	historyDir      string
}

func (u *userData) makeDirectories() {
	// We ignore errors, as we should function fine without these
	_ = os.MkdirAll(u.configDir, 0o700)
	_ = os.MkdirAll(u.systemConfigDir, 0o755)
	_ = os.MkdirAll(u.dataDir, 0o700)
	_ = os.MkdirAll(u.cacheDir, 0o700)
	_ = os.MkdirAll(u.historyDir, 0o700)
}

var rootCmd = &cobra.Command{
	Use: "herd",
	Long: `Replace your ssh for loops with a tool that
- can find hosts for you
- can handle thousands of hosts in parallel
- does not fork a command for every host
- stores all its history, including output, for you to reuse
- can run interactively!
`,
	Example: `  herd run '*' os=Debian -- dpkg -l bash
  herd interactive *vpn-gateway*`,
	Args:    cobra.NoArgs,
	Version: herd.Version(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if f := viper.GetString("Profile"); f != "" {
			pfd, err := os.Create(f + ".cpuprofile")
			if err != nil {
				bail("Could not create CPU profile file: ", err)
			}
			if err = pprof.StartCPUProfile(pfd); err != nil {
				bail("Could not start CPU profile: ", err)
			}
			tfd, err := os.Create(f + ".trace")
			if err != nil {
				bail("Could not create trace file: ", err)
			}
			if err = trace.Start(tfd); err != nil {
				bail("Could not start trace: ", err)
			}
			cmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
				pprof.StopCPUProfile()
				pfd.Close()
				trace.Stop()
				tfd.Close()
			}
		}
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var err error
	currentUser, err = getCurrentUser()
	if err != nil {
		bail("%s", err)
	}
	currentUser.makeDirectories()
	rootCmd.SetHelpTemplate(fmt.Sprintf(`%s

Configuration: %s, %s
Datadir: %s
History: %s
Cache: %s
Providers: %s
`,
		rootCmd.HelpTemplate(),
		filepath.Join(currentUser.configDir, "config.yaml"),
		filepath.Join(currentUser.systemConfigDir, "config.yaml"),
		currentUser.dataDir,
		currentUser.historyDir,
		currentUser.cacheDir,
		strings.Join(herd.Providers(), ",")))
	defaultAgentTimeout := 50 * time.Millisecond
	if _, ok := os.LookupEnv("SSH_CONNECTION"); ok {
		defaultAgentTimeout = 200 * time.Millisecond
	}
	cobra.OnInitialize(initConfig)
	f := rootCmd.PersistentFlags()
	f.Duration("splay", 0, "Wait a random duration up to this argument before and between each host")
	f.DurationP("timeout", "t", 60*time.Second, "Global timeout for commands")
	f.Duration("load-timeout", 30*time.Second, "Timeout for loading host data from providers")
	f.Duration("host-timeout", 10*time.Second, "Per-host timeout for commands")
	f.Duration("connect-timeout", 3*time.Second, "Per-host ssh connect timeout")
	f.Duration("ssh-agent-timeout", defaultAgentTimeout, "SSH agent timeout when checking functionality")
	f.IntP("parallel", "p", 0, "Maximum number of hosts to run on in parallel")
	f.StringP("output", "o", "all", "When to print command output (all at once, per host or per line)")
	f.Bool("no-pager", false, "Disable the use of the pager")
	f.Bool("no-color", false, "Disable the use of the colors in the output")
	f.StringP("loglevel", "l", "INFO", "Log level")
	f.StringSliceP("sort", "s", []string{"name"}, "Sort hosts by these fields before running commands")
	f.Bool("timestamp", false, "In tail mode, prefix each line with the current time")
	f.String("profile", "", "Write profiling and tracing data to files starting with this name")
	f.Bool("refresh", false, "Force caches to be refreshed")
	f.Bool("no-refresh", false, "Do not try to refresh cached data")
	f.Bool("strict-loading", false, "Fail if any provider fails to load data")
	f.Bool("no-magic-providers", false, "Do not use magic autodiscovery, only explicitly configured providers")
	bindFlagsAndEnv(f)
}

func bindFlagsAndEnv(s *pflag.FlagSet) {
	rx := regexp.MustCompile("((?:^|-).)")
	toUpper := func(s string) string {
		return strings.ToUpper(strings.Trim(s, "-"))
	}
	s.VisitAll(func(f *pflag.Flag) {
		varName := rx.ReplaceAllStringFunc(f.Name, toUpper)
		envName := "HERD_" + strings.ReplaceAll(strings.ToUpper(f.Name), "-", "_")
		if err := viper.BindPFlag(varName, f); err != nil {
			panic(err)
		}
		if err := viper.BindEnv(varName, envName); err != nil {
			panic(err)
		}
	})
}

func initConfig() {
	viper.AddConfigPath(currentUser.configDir)
	viper.AddConfigPath("/etc/herd")
	viper.SetConfigName("config")

	// Read the configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			bail("Can't read configuration: %s", err)
		}
	}

	// Check configuration variables
	level, err := logrus.ParseLevel(viper.GetString("LogLevel"))
	if err != nil {
		bail("Unknown loglevel: %s. Known loglevels: DEBUG, INFO, NORMAL, WARNING, ERROR", viper.GetString("LogLevel"))
	}
	logrus.SetLevel(level)

	if viper.GetBool("NoColor") {
		ansi.DisableColors(true)
	}
	outputModes := map[string]herd.OutputMode{
		"all":      herd.OutputAll,
		"inline":   herd.OutputInline,
		"per-host": herd.OutputPerhost,
		"tail":     herd.OutputTail,
	}
	om, ok := outputModes[viper.GetString("Output")]
	if !ok {
		bail("Unknown output mode: %s. Known modes: all, inline, per-host, tail", viper.GetString("Output"))
	}
	viper.Set("Output", om)
}

func bail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func setupScriptEngine(executor herd.Executor) (*scripting.ScriptEngine, error) {
	hosts := new(herd.HostSet)
	hosts.SetSortFields(viper.GetStringSlice("Sort"))
	ui := herd.NewSimpleUI(hosts)
	ui.SetOutputMode(viper.Get("Output").(herd.OutputMode))
	ui.SetOutputTimestamp(viper.GetBool("Timestamp"))
	ui.SetPagerEnabled(!viper.GetBool("NoPager"))
	ui.BindLogrus()

	registry := herd.NewRegistry(currentUser.dataDir, currentUser.cacheDir)
	conf := viper.Sub("Providers")
	if conf != nil {
		if err := registry.LoadProviders(conf); err != nil {
			logrus.Error(err.Error())
			ui.End()
			return nil, err
		}
	}
	if !viper.GetBool("NoMagicProviders") {
		registry.LoadMagicProviders()
	}
	if viper.GetBool("Refresh") {
		registry.InvalidateCache()
	}
	if viper.GetBool("NoRefresh") {
		registry.KeepCaches()
	}
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("LoadTimeout"))
	defer cancel()
	if err := registry.LoadHosts(ctx, ui.LoadingMessage); err != nil {
		// Do not log this error, registry.LoadHosts() does its own error logging
		if viper.GetBool("StrictLoading") {
			ui.End()
			return nil, err
		}
	}
	if err := registry.LoadHostKeys(ctx, ui.LoadingMessage); err != nil {
		if viper.GetBool("StrictLoading") {
			ui.End()
			return nil, err
		}
	}
	ui.Sync()
	runner := herd.NewRunner(hosts, executor)
	handleSignals(runner)
	runner.SetSplay(viper.GetDuration("Splay"))
	runner.SetParallel(viper.GetInt("Parallel"))
	runner.SetTimeout(viper.GetDuration("Timeout"))
	runner.SetHostTimeout(viper.GetDuration("HostTimeout"))
	runner.SetConnectTimeout(viper.GetDuration("ConnectTimeout"))
	return scripting.NewScriptEngine(hosts, ui, registry, runner), nil
}
