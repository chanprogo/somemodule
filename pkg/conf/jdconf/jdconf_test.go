package jdconf

import (
	"runtime"
	"testing"
)

var cfg Config

func TestTwo(t *testing.T) {

	err := readConfig(&cfg)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("Running with %v threads\n", cfg.Threads)
	t.Logf("Name %v \n", cfg.Name)
	t.Logf("Version %v \n", cfg.Version)

	if cfg.Threads > 0 {
		runtime.GOMAXPROCS(cfg.Threads)
		t.Logf("Running with %v threads", cfg.Threads)
	}

	// quit := make(chan bool)
	// <-quit
}
