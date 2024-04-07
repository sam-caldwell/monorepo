package application

import "github.com/sam-caldwell/monorepo/go/ansi"

// Server - Runs a listener/event collector service
func (app *Application) Server() (err error) {
	ansi.Green().Println("server starting...")
	return app.server.Listener().Error()
}
