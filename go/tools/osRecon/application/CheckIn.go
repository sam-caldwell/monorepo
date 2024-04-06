package application

// CheckIn - Performs network request (healthcheck) which also polls server for pending queries to execute
func (app *Application) CheckIn() error {

	// This go routine will call to the server and both perform a healthcheck
	// and receive any waiting queries the server wishes the client to execute.
	// These queries are received, sanitized and pushed to queryQueue.

	// start the go routine

	return nil
}
