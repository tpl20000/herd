package cmd

import (
	"fmt"
	"os"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/seveas/herd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	Version: "1.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		herd.UI = herd.NewSimpleUI()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().DurationP("timeout", "t", 60*time.Second, "Global timeout for commands")
	rootCmd.PersistentFlags().Duration("host-timeout", 10*time.Second, "Per-host timeout for commands")
	rootCmd.PersistentFlags().Duration("connect-timeout", 3*time.Second, "Per-host ssh connect timeout")
	rootCmd.PersistentFlags().IntP("parallel", "p", 0, "Maximum number of hosts to run on in parallel")
	rootCmd.PersistentFlags().StringP("output", "o", "all", "When to print command output (all at once, per host or per line)")
	viper.BindPFlag("Timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	viper.BindPFlag("HostTimeout", rootCmd.PersistentFlags().Lookup("host-timeout"))
	viper.BindPFlag("ConnectTimeout", rootCmd.PersistentFlags().Lookup("connect-timeout"))
	viper.BindPFlag("Parallel", rootCmd.PersistentFlags().Lookup("parallel"))
	viper.BindPFlag("Output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	viper.SetDefault("HistoryDir", path.Join(home, ".herd", "history"))
	viper.SetDefault("LogLevel", herd.INFO)
	viper.SetDefault("Formatter", "pretty")

	viper.AddConfigPath(path.Join(home, ".herd"))
	viper.AddConfigPath("/etc/herd")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintln(os.Stderr, "Can't read configuration:", err)
			os.Exit(1)
		}
	}

	viper.SetEnvPrefix("herd")
	viper.AutomaticEnv()

	// Check configuration variables
	if _, ok := herd.Formatters[viper.GetString("Formatter")]; !ok {
		fmt.Fprintln(os.Stderr, "Unknown formatter:", viper.GetString("Formatter"))
		os.Exit(1)
	}
	logLevels := map[string]int{"DEBUG": herd.DEBUG, "INFO": herd.INFO, "NORMAL": herd.NORMAL, "WARNING": herd.WARNING, "ERROR": herd.ERROR}
	if level, ok := logLevels[viper.GetString("LogLevel")]; ok {
		viper.Set("LogLevel", level)
	}
}
