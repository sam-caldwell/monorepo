package application

func (app *Application) Emitter() error {
	// The Emitter will receive message structures from eventQueue
	// Each eventQueue message will be sent over the network to the server.
	return nil
}
