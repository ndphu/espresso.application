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
	StartupCalled  bool
	RunCalled      bool
	ShutdownCalled bool
}

func (a *ApplicationTest) Startup() error {
	a.StartupCalled = true
	return nil
}

func (a *ApplicationTest) Run() {
	a.RunCalled = true
}

func (a *ApplicationTest) Shutdown() error {
	a.ShutdownCalled = true
	return nil
}

func TestApplication(t *testing.T) {
	testApp := new(ApplicationTest)
	//testApp.SignalChannel = make(chan os.Signal, 1)
	am := application.NewApplicationManager()

	go func() {
		time.Sleep(200 * time.Millisecond)
		//close(testApp.GetSignalChannel())
		am.StopApplication()
	}()

	am.RunApplication(testApp)

	assert.True(t, testApp.StartupCalled, "Startup Callback should be called")
	assert.True(t, testApp.RunCalled, "Run Callback should be called")
	assert.True(t, testApp.ShutdownCalled, "Shutdown Callback should be called")
}
