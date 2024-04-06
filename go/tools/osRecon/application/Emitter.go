package application

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"time"
)

// Emitter - Start the go routine which will emit messages to the server
func (app *Application) Emitter() error {
	//ToDo: initialize emitter stats
	go func() {
		//disconnect from the server on termination.
		defer app.server.Disconnect()
		//Loop indefinitely to poll the
		for {
			//ToDo: update emitter stats
			event := <-app.eventQueue
			event.SendTime = time.Now()
			//connect (if not connected) and send the event.
			if err := app.server.Connect().SendEvent(event).Error(); err != nil {
				ansi.Red().Printf("Emitter error: %v", err).LF().Reset()
				//ToDo: update emitter stats
			}
			//ToDo: update emitter stats
		}
	}()
	return nil
}
