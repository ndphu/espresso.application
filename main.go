package application

import (
	"os"
	"os/signal"
	"syscall"
)

type ApplicationManager struct {
	SignalChannel chan os.Signal
}

type Application interface {
	Startup() error
	Run()
	Shutdown() error
}

type Callback func()

func NewApplicationManager() *ApplicationManager {
	am := ApplicationManager{}
	am.SignalChannel = make(chan os.Signal, 1)
	return &am
}

func (am *ApplicationManager) StopApplication() {
	close(am.SignalChannel)
}

func (am *ApplicationManager) RunApplication(a Application) error {
	err := a.Startup()
	if err != nil {
		return err
	}
	go a.Run()
	//sigchan := a.GetSignalChannel()
	sigchan := am.SignalChannel
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
	return a.Shutdown()
}
