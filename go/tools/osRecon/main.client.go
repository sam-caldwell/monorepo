package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/client"
	"os"
	"path/filepath"
)

const (
	clientUsage = `
(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>

This is the osRecon client which queries a local host where it is running and emits the collected
telemetry to the osRecon server.

Syntax:
    %s --host <hostname> --port <portNumber> --apiKey <base64 key string> [optional flags]

Optional Flags:
    --eventQueueSz <int> : size of the event queue used by the collection and emitter
    --queryQueueSz <int> : size of the query queue used to fetch and stage queries for execution.
    --queryQueuePollInterval <int> : number of seconds (approx.) between client to server requests for more queries

Options:
    -h | --help displays this message
`
)

func main() {
	programName := filepath.Base(os.Args[0])
	cli.GetHelp(fmt.Sprintf(clientUsage, programName))

	ansi.Blue().Println("Starting...").Reset()
	var app client.Client
	exit.TerminateOnError(app.Configure())
	exit.TerminateOnError(app.Emitter())
	exit.TerminateOnError(app.CheckIn())
	exit.TerminateOnError(app.ProcMon())
	exit.TerminateOnError(app.NetMon())
	exit.TerminateOnError(app.HwMon())
	exit.TerminateOnError(app.FsMon())
	exit.TerminateOnError(app.SysMon())
	exit.TerminateOnError(app.Query())
	exit.TerminateOnError(app.SignalHandler())
	ansi.Green().Println("Terminating").Reset()
}
