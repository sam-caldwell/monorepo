package client

import (
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/client/connector"
	"github.com/sam-caldwell/monorepo/go/types"
)

type Client struct {
	// This is the configuration and message queues (channels)
	// shared between the asynchronous processes
	Mode types.OperatingMode

	// Message Queue Channels
	eventQueue chan types.Event
	queryQueue chan types.ThreatQlQuery

	//How often do we poll the server for queries to run
	queryQueuePollInterval uint16

	//The client's server connection information
	server connector.Server
}
