package smtp

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/secrets/volatileSecrets"
)

type Client struct {
	host     net.Fqdn
	port     net.PortNumber
	userName string
	password volatileSecrets.Password
}
