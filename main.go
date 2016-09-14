package application

import (
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	StartupCallback  Callback
	RunCallback      Callback
	ShutdownCallback Callback
	sigChannel       chan os.Signal
}

type Callback func()

func New(startupCallback Callback, runCallback Callback, shudownCallback Callback) *Application {
	app := new(Application)
	// Set callbacks
	app.StartupCallback = startupCallback
	app.RunCallback = runCallback
	app.ShutdownCallback = shudownCallback
	// Create signal channel
	app.sigChannel = make(chan os.Signal, 1)
	signal.Notify(app.sigChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	return app
}

func (a *Application) Start() {
	a.StartupCallback()
	// Handle signal

	go a.RunCallback()

	for {
		sig, success := <-a.sigChannel
		if !success ||
			sig == syscall.SIGINT ||
			sig == syscall.SIGTERM ||
			sig == syscall.SIGKILL ||
			sig == os.Interrupt ||
			sig == os.Kill {
			break
		}
	}
	a.ShutdownCallback()
}

func (a *Application) Stop() {
	close(a.sigChannel)
}
