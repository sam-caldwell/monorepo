package application

import (
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server"
	"github.com/sam-caldwell/monorepo/go/types"
)

type Application struct {
	// This is the configuration and message queues (channels)
	// shared between the asynchronous processes
	Mode types.OperatingMode

	// Message Queue Channels
	eventQueue chan types.Event
	queryQueue chan types.ThreatQlQuery

	//Queue sizes (default to constant, override by cli arg)
	//Size in number of records per queue
	eventQueueSz uint16
	queryQueueSz uint16

	//How often do we poll the server for queries to run
	queryQueuePollInterval uint16

	//Server configuration
	server server.Server
}
