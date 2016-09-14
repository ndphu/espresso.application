package application_test

import (
	application "github.com/ndphu/espresso.application"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApplication(t *testing.T) {
	startupCallbackCalled := false
	shutdownCallbackCalled := false
	runCallbackCalled := false
	app := application.New(
		func() {
			t.Log("Startup Callback is called")
			startupCallbackCalled = true
		},
		func() {
			t.Log("Run Callback is called")
			runCallbackCalled = true
		},
		func() {
			t.Log("Shutdown Callback is called.")
			shutdownCallbackCalled = true
		})
	go func() {
		time.Sleep(500 * time.Millisecond)
		app.Stop()
	}()
	app.Start()
	assert.True(t, startupCallbackCalled, "Startup Callback should be called")
	assert.True(t, shutdownCallbackCalled, "Shutdown Callback should be called")
}
