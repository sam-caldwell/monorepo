package connector

import (
	"github.com/sam-caldwell/monorepo/projects/net"
	"github.com/sam-caldwell/monorepo/projects/types"
)

// Connector - This is the application server connection information.
type Connector struct {
	host   net.Fqdn
	port   net.PortNumber
	apiKey types.ApiKey
	err    error
}
