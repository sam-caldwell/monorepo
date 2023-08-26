package main

import (
	ansi2 "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
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

	ansi2.Blue().Println("init-repo: starting")

	var run runcommand.Runner
	err := run.Run("go mod tidy").
		Run("go install honnef.co/go/tools/cmd/staticcheck@latest").
		Run("pip3 install -r ./requirements.txt").
		Run("npm install").
		Run("node --version").
		Error()
	ansi2.Blue().Println("Recovering")
	out := run.Output()
	if strings.TrimSpace(out) == "" {
		out = "<no output>"
	}

	if err != nil {
		ansi2.Red().
			Printf("Error initializing repo: %v\n%s", err, out).
			LF().
			Reset().
			Fatal(exit.GeneralError)
	}
	ansi2.Green().
		Printf("Success: '%s'", out).
		LF().
		Reset()
}
