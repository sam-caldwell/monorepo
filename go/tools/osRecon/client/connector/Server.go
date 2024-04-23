package connector

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Connector - This is the application server connection information.
type Connector struct {
	host   net.Fqdn
	port   net.PortNumber
	apiKey types.ApiKey
	err    error
}
