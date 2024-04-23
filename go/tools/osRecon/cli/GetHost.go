package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetHost - Return the value from --host
func GetHost() net.Fqdn {

	argServerHost, err := simpleArgs.GetOptionValue("--host")

	if err != nil {

		ansi.Red().
			Printf("Parsing error (%s): %v", "--host", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()

	}

	var serverHost net.Fqdn

	exit.TerminateOnError(serverHost.FromString(&argServerHost))

	return serverHost

}
