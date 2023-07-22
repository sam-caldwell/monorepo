package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

const (
	commandUsage = `
lint-repo -h|--help 
   Show usage
   
lint-repo [-color]
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
