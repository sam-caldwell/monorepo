package application

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/types"
	"time"
)

// Emitter - Start the go routine which will emit messages to the server
//
//	The emitter will consume messages from eventQueue and send them
//	to the server.
func (app *Application) Emitter() (err error) {
	ansi.Green().Println("starting Emitter()")
	//ToDo: initialize emitter stats

	//Make sure we can connect to the server and send an event or return an error
	err = app.server.Connect().SendEvent(types.Event{
		EventTime: time.Now(),
		SendTime:  time.Now(),
		Resource:  types.ResourceId(0),
		Field:     types.FieldId(0),
		Value:     []byte{0x00}, //Todo: we need a formal starting event
	}).Error()

	if err == nil {
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
	} else {
		ansi.Red().Printf("Emitter failed to connect to server. %v", err).LF().Reset()
	}
	return err
}
