package application

import (
	"os"
	"os/signal"
	"syscall"
)

type Application interface {
	Startup()
	GetSignalChannel() chan os.Signal
	Run()
	Shutdown()
}

func RunApplication(a Application) {
	a.Startup()
	go a.Run()
	sigchan := a.GetSignalChannel()
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	for {
		sig, success := <-sigchan
		if !success ||
			sig == syscall.SIGINT ||
			sig == syscall.SIGTERM ||
			sig == syscall.SIGKILL ||
			sig == os.Interrupt ||
			sig == os.Kill {
			break
		}
	}
	a.Shutdown()
}
