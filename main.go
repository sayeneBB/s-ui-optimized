package main

import (
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/admin8800/s-ui/app"
	"github.com/admin8800/s-ui/cmd"
)

func runApp() {
	// Set Go memory soft limit to 384MB — prevents unbounded heap growth
	// that triggers swap on small VPSs. Go may briefly exceed this but GC
	// will bring it back under control. Adjust via GOMEMLIMIT env var.
	debug.SetMemoryLimit(384 * 1024 * 1024)
	app := app.NewApp()

	err := app.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	// Trap shutdown signals
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM)
	for {
		sig := <-sigCh

		switch sig {
		case syscall.SIGHUP:
			app.RestartApp()
		default:
			app.Stop()
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		runApp()
		return
	} else {
		cmd.ParseCmd()
	}
}
