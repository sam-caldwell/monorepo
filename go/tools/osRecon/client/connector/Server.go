package connector

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Server - This is the application server connection information.
type Server struct {
	host   net.Fqdn
	port   net.PortNumber
	apiKey types.ApiKey
	err    error
}
