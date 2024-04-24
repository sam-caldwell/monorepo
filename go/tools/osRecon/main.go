package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/client"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
	"path/filepath"
)

const (
	bannerText   = "(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>"
	commandUsage = `
%s

Syntax:
    %s [mode]

modes:
    Run as client mode:

        <cmd> client --host <hostname> --port <portNumber> --apiKey <base64 key string>

    Run as server mode:

        <cmd> server --host <hostname> --port <portNumber> --apiKey <base64 key string>

Options:
    -h | --help displays this message
`
)

func main() {
	programName := filepath.Base(os.Args[0])

	cli.GetHelp(fmt.Sprintf(commandUsage, bannerText, programName))

	ansi.Blue().Println("Starting...").Reset()

	mode, err := cli.GetMode()
	exit.TerminateOnError(err)

	switch mode {
	case types.Client:
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
	case types.Server:
		var app server.Server
		exit.TerminateOnError(app.Configure())
		exit.TerminateOnError(app.Listener())
		exit.TerminateOnError(app.SignalHandler())

	default:
		ansi.Red().Println("Error: Invalid Mode").Fatal(exit.GeneralError).Reset()
	}
	ansi.Green().Println("Terminating").Reset()
}
