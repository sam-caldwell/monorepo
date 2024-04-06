package server

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/types"
)

type Server struct {
	host   net.Fqdn
	port   net.PortNumber
	apiKey types.ApiKey
	//Todo: add some stats
}
