package smtp

import (
	"github.com/sam-caldwell/monorepo/projects/net"
	"github.com/sam-caldwell/monorepo/projects/secrets/volatileSecrets"
)

type Client struct {
	host     net.Fqdn
	port     net.PortNumber
	userName string
	password volatileSecrets.Password
}
