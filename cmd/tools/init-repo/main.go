package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	"strings"
)

const (
	commandUsage = `
init-repo -h|-help
	show this screen

run-builds
	Run the builds for every command project in the repo.

 * Supported on--
		* ubuntu linux
		* windows
		* macos
`
)

func main() {

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	ansi.Blue().Println("init-repo: starting")

	var run runcommand.Runner
	err := run.Run("go mod tidy").
		Run("go install honnef.co/go/tools/cmd/staticcheck@latest").
		Run("pip3 install -r ./requirements.txt").
		Run("npm install").
		Error()
	ansi.Blue().Println("Recovering")
	out := run.Output()
	if strings.TrimSpace(out) == "" {
		out = "<no output>"
	}

	if err != nil {
		ansi.
			Red().
			Printf("Error initializing repo: %v\n%s", err, out).
			LF().
			Reset().
			Fatal(exit.GeneralError)
	}
	ansi.
		Green().
		Printf("Success: '%s'", out).
		LF().
		Reset()
}
