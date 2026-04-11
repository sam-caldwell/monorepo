package connector

import (
	"github.com/sam-caldwell/monorepo/projects/net"
	"github.com/sam-caldwell/monorepo/projects/types"
)

// Configure - configure the client connector properties.
func (svr *Connector) Configure(host net.Fqdn, port net.PortNumber, apiKey types.ApiKey) {
	svr.host = host
	svr.port = port
	svr.apiKey = apiKey
}
