package application_test

import (
	application "github.com/ndphu/espresso.application"
	"github.com/stretchr/testify/assert"
	//"os"
	"testing"
	"time"
)

type ApplicationTest struct {
	//SignalChannel  chan os.Signal
	StopRequestCallback application.Callback
	StartupCalled       bool
	RunCalled           bool
	ShutdownCalled      bool
}

func (a *ApplicationTest) Startup() {
	a.StartupCalled = true
}

func (a *ApplicationTest) Run() {
	a.RunCalled = true
}

func (a *ApplicationTest) Shutdown() {
	a.ShutdownCalled = true
}

// func (a *ApplicationTest) GetSignalChannel() chan os.Signal {
// 	return a.SignalChannel
// }

func (a *ApplicationTest) SetStopRequestCallback(callback application.Callback) {
	a.StopRequestCallback = callback
}

func TestApplication(t *testing.T) {
	testApp := new(ApplicationTest)
	//testApp.SignalChannel = make(chan os.Signal, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		//close(testApp.GetSignalChannel())
		testApp.StopRequestCallback()
	}()

	application.RunApplication(testApp)

	assert.True(t, testApp.StartupCalled, "Startup Callback should be called")
	assert.True(t, testApp.RunCalled, "Run Callback should be called")
	assert.True(t, testApp.ShutdownCalled, "Shutdown Callback should be called")
}
