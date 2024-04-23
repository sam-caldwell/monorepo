package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetPort - return the port number (uint16) from --port
func GetPort() net.PortNumber {

	argServerPort, err := simpleArgs.GetOptionIntValue("--port", true)

	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--port", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	var serverPort net.PortNumber

	exit.TerminateOnError(serverPort.FromInt(argServerPort))

	return serverPort

}
