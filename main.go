package application

import (
	"os"
	"os/signal"
	"syscall"
)

type Application interface {
	Startup()
	SetStopRequestCallback(stopCallback Callback)
	Run()
	Shutdown()
}

type Callback func()

func RunApplication(a Application) {
	a.Startup()
	go a.Run()
	//sigchan := a.GetSignalChannel()
	sigchan := make(chan os.Signal, 1)
	a.SetStopRequestCallback(func() {
		close(sigchan)
	})
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
