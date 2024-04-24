package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetHost - Return the value from --host
func GetHost() net.Fqdn {
	const argString = "--host"
	argServerHost, err := simpleArgs.GetOptionValue(argString)

	if err != nil {

		ansi.Red().
			Printf("Parsing error (%s): %v", argString, err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()

	}

	var serverHost net.Fqdn

	exit.TerminateOnError(serverHost.FromString(&argServerHost))

	return serverHost

}
