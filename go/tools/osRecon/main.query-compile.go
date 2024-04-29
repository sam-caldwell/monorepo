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
	queryCompileUsage = `
(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>

Given a query on stdin, compile the query into its binary file form and
store it in the queryDir directory location for consumption by osRecon server.

Syntax:
    %s --queryDir <path>

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
