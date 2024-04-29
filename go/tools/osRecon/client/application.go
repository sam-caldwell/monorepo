package client

import (
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/client/connector"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/threatQL"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Client - osRecon Client
type Client struct {
	// This is the configuration and message queues (channels)
	// shared between the asynchronous processes
	Mode types.OperatingMode

	// Message Queue Channels
	eventQueue chan types.Event
	queryQueue chan threatQL.Query

	//How often do we poll the server for queries to run
	queryQueuePollInterval uint16

	//The client's server connection information
	server connector.Connector
}
