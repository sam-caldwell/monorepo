package client

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"time"
)

// CheckIn - Performs network request (healthcheck) which also polls server for pending queries to execute
func (app *Application) CheckIn() error {
	ansi.Green().Println("starting CheckIn()")
	go func() {
		// This go routine will call to the server and both perform a healthcheck
		// and receive any waiting queries the server wishes the client to execute.
		// These queries are received, sanitized and pushed to queryQueue.
		for i := 0; ; i++ {
			//ToDo: update stats
			if err := app.server.PollQueue(app.queryQueue).Error(); err != nil {
				//ToDo: update stats
				ansi.Red().Printf("Server checkin failed: %v\n", err).Reset()
			}
			//ToDo: update stats
			time.Sleep(time.Duration(app.queryQueuePollInterval) * time.Second)
			ansi.Green().Printf("server checkin completed %d\n", i).Reset()
		}
	}()
	return nil
}
