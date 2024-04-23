package client

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/types"
)

func (app *Client) Configure() error {

	var err error
	var eventQueueSz uint16
	var queryQueueSz uint16

	ansi.Green().Println("configuring application...")
	//
	// Process command-line args
	//
	// Get the mode (client, server)
	//
	app.Mode, err = cli.GetMode()
	if err != nil {
		return err
	}
	//
	// Get server Host, port, ApiKey
	//
	app.server.Host = cli.GetHost()
	app.server.Port = cli.GetPort()
	app.server.ApiKey = cli.GetApiKey()
	//
	// Get the queue sizes
	//
	if eventQueueSz, err = cli.GetEventQueueSz(defaultEventQueueSz); err != nil {
		return err
	}
	if queryQueueSz, err = cli.GetQueryQueueSz(defaultQueryQueueSz); err != nil {
		return err
	}
	if app.queryQueuePollInterval, err = cli.GetQueryQueuePollInterval(defaultQueryQueuePollInterval); err != nil {
		return err
	}
	// setup channel for collecting events from monitors (eventQueue)
	app.eventQueue = make(chan types.Event, eventQueueSz)
	// setup channel for receiving queries from the server (queryQueue)
	app.queryQueue = make(chan types.ThreatQlQuery, queryQueueSz)

	return nil
}
