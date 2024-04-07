package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/application"
	"github.com/sam-caldwell/monorepo/go/types"
)

func main() {

	var app application.Application

	ansi.Blue().Println("Starting...").Reset()

	exit.TerminateOnError(app.Configure())

	switch app.Mode {
	case types.Client:
		exit.TerminateOnError(app.Emitter())
		exit.TerminateOnError(app.CheckIn())
		exit.TerminateOnError(app.ProcMon())
		exit.TerminateOnError(app.NetMon())
		exit.TerminateOnError(app.HwMon())
		exit.TerminateOnError(app.FsMon())
		exit.TerminateOnError(app.SysMon())
		exit.TerminateOnError(app.Query())
	case types.Server:
		exit.TerminateOnError(app.Server())
	default:
		ansi.Red().Println("Error: Invalid Mode").Fatal(exit.GeneralError).Reset()
	}

	exit.TerminateOnError(app.SignalHandler())

	ansi.Green().Println("Terminating").Reset()

}
