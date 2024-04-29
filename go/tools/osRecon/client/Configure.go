package client

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/threatQL"
	"github.com/sam-caldwell/monorepo/go/types"
)

func (app *Client) Configure() error {

	var err error
	var eventQueueSz uint16
	var queryQueueSz uint16

	ansi.Green().Println("configuring application...")
	//
	// Get server connection Host, port, ApiKey
	//
	app.server.Configure(cli.GetHost(), cli.GetPort(), cli.GetApiKey())
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
	app.queryQueue = make(chan threatQL.Query, queryQueueSz)

	return nil
}
