package main

import (
	ansi "github.com/sam-caldwell/monorepo/v2/projects/go/ansi"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit"
	"github.com/sam-caldwell/monorepo/v2/projects/go/runcommand"
)

const (
	commandUsage = `
run-linter -h|--help 
   Show usage
   
run-linter [-color]
   Run linter and if -color is present show
   the output in ANSI color.
`
	displayWidth = 40
)

func main() {

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	if out, err := runcommand.ShellExecute("snyk test --file=go.mod"); err != nil {
		ansi.Red().Printf("GO: snyk FAIL: %s\nout: %s\n", err, out).Reset().Fatal(exit.GeneralError)
	}
	ansi.Green().Println("GO: snyk PASS").Reset()

	if out, err := runcommand.ShellExecute("snyk test --file=package.json"); err != nil {
		ansi.Red().Printf("NODE: snyk FAIL: %s\nout: %s\n", err, out).Reset().Fatal(exit.GeneralError)
	}
	ansi.Green().Println("NODE: snyk PASS").Reset()
}
