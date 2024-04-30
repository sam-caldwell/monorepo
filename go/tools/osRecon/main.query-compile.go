package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/threatQL"
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
	cli.GetHelp(fmt.Sprintf(queryCompileUsage, programName))
	var app threatQL.Compiler
	exit.TerminateOnError(app.Configure())
	exit.TerminateOnError(app.Compile())
	exit.TerminateOnError(app.SignalHandler())
	ansi.Blue().Println("Starting...").Reset()
	//Todo: implement query-compile
	ansi.Green().Println("Terminating").Reset()
}
