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
	analyzerUsage = `
(c) 2024 Sam Caldwell.  <mail@samcaldwell.net>

This is the osRecon analyzer which evaluates telemetry from osRecon server for indicators of compromise.

Syntax:
    %s --queryDir <path> --statsDir <path> --alertAddress <email address>

Options:
    -h | --help displays this message
`
)

func main() {
	programName := filepath.Base(os.Args[0])
	cli.GetHelp(fmt.Sprintf(analyzerUsage, programName))

	ansi.Blue().Println("Starting...").Reset()

	var app server.Server
	exit.TerminateOnError(app.Configure())
	exit.TerminateOnError(app.Listener())
	exit.TerminateOnError(app.SignalHandler())
}
