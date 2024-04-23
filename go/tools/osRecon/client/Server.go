package client

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Server - This is the application server connection information.
type Server struct {
	Host   net.Fqdn
	Port   net.PortNumber
	ApiKey types.ApiKey
}
