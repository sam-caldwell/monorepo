package application

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
	"path/filepath"
)

func (app *Application) Configure() error {
	bannerText := "(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>"
	programName := filepath.Base(os.Args[0])
	// Process command-line args

	argMode := simpleArgs.GetCommand(fmt.Sprintf(commandUsage, bannerText, programName))
	argServerHost, err := simpleArgs.GetOptionValue("--host")
	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--host", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	argServerPort, err := simpleArgs.GetOptionIntValue("--port", true)
	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--port", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	argServerApiKey, err := simpleArgs.GetOptionValue("--apiKey")
	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--apiKey", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	app.eventQueueSz, err = simpleArgs.GetOptionUint16Value("--eventQueueSz", false)
	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", "--eventQueueSz", err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	app.queryQueueSz, err = simpleArgs.GetOptionUint16Value("--queryQueueSz", false)
	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", "--queryQueueSz", err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	if app.queryQueueSz == 0 {
		app.queryQueueSz = defaultQueryQueueSz
	}
	if app.eventQueueSz == 0 {
		app.eventQueueSz = defaultEventQueueSz
	}
	var serverHost net.Fqdn
	var serverPort net.PortNumber
	var serverApiKey types.ApiKey

	exit.TerminateOnError(app.Mode.FromString(argMode))
	exit.TerminateOnError(serverHost.FromString(&argServerHost))
	exit.TerminateOnError(serverPort.FromInt(argServerPort))
	exit.TerminateOnError(serverApiKey.FromString(&argServerApiKey))
	app.server = *server.New(serverHost, serverPort, serverApiKey)

	// setup channel for collecting events from monitors (eventQueue)
	app.eventQueue = make(chan types.Event, app.eventQueueSz)
	// setup channel for receiving queries from the server (queryQueue)
	app.queryQueue = make(chan types.ThreatQlQuery, app.queryQueueSz)
	return nil
}
