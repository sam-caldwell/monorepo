package server

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server/persistence"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Server - osRecon Server
type Server struct {
	//
	// Server connection info...
	//
	host   net.Fqdn       //the server's hostname
	port   net.PortNumber // the tcp port number used for the server connect
	apiKey types.ApiKey   // The auth key used to connect ot the server
	//
	// Internal state
	//
	//ToDo: Add some stats...
	//
	// The event collector (queue and disk writer)
	//
	eventCollector persistence.EventCollector
	//
	// QueryQueue
	//
	queryQueue persistence.QueryCollector
}
