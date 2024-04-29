package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server"
	"os"
	"path/filepath"
)

const (
	serverUsage = `
(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>

This is the osRecon server which receives all telemetry from the osRecon client.

Syntax:
    %s --host <hostname> --port <portNumber> --apiKey <base64 key string> --queryDir <path> --statsDir <path>

Options:
    -h | --help displays this message
`
)

func main() {
	programName := filepath.Base(os.Args[0])
	cli.GetHelp(fmt.Sprintf(serverUsage, programName))

	ansi.Blue().Println("Starting...").Reset()

	var app server.Server
	exit.TerminateOnError(app.Configure())
	exit.TerminateOnError(app.Listener())
	exit.TerminateOnError(app.SignalHandler())
}
