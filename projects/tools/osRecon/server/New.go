package server

import (
	"github.com/sam-caldwell/monorepo/projects/net"
	"github.com/sam-caldwell/monorepo/projects/types"
)

// New - Initialize and return a new Server object
func New(host net.Fqdn, port net.PortNumber, apiKey types.ApiKey) *Server {
	return &Server{
		host:   host,
		port:   port,
		apiKey: apiKey,
	}
}
