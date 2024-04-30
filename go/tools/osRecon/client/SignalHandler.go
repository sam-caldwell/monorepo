package client

import (
	"os"
	"os/signal"
)

// SignalHandler - Client signal handler to grab signals and terminate gracefully
func (app *Client) SignalHandler() error {
	// This method handles signals (SIGINT) from the operating system
	// and cleanly terminates the system.

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c
	//Terminate gracefully
	//ToDo: we should allow all queries to be run
	//ToDo: we should allow all events to be emitted
	return nil
}
