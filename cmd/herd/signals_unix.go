//go:build unix
// +build unix

package main

import (
	"os"
	"syscall"

	"github.com/seveas/herd"

	"github.com/sirupsen/logrus"
)

func handleSignals(r *herd.Runner) {
	r.OnSignal(os.Interrupt, func() {
		logrus.Errorf("Interrupted, canceling with unfinished tasks")
		r.Interrupt()
	})
	r.OnSignal(syscall.SIGUSR1, func() {
		_, s := r.Settings()
		p := s["Parallel"].(int) * 3 / 2
		logrus.Infof("Increasing parallelism to %d", p)
		r.SetParallel(p)
	})
	r.OnSignal(syscall.SIGUSR2, func() {
		_, s := r.Settings()
		p := s["Parallel"].(int) / 2
		logrus.Infof("Decreasing parallelism to %d", p)
		r.SetParallel(p)
	})
}
