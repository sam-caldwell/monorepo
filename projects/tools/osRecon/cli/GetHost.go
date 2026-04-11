package cli

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/net"
	"github.com/sam-caldwell/monorepo/projects/simpleArgs"
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
