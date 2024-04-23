package server

import (
	"os"
	"os/signal"
)

func (svr *Server) SignalHandler() error {
	// This method handles signals (SIGINT) from the operating system
	// and cleanly terminates the system.

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c
	return nil
}
