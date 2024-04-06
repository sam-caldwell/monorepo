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

	//Server configuration
	server server.Server
}
