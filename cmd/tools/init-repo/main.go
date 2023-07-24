package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repotools "github.com/sam-caldwell/go/v2/projects/repotools/init-repo"
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

	commands := []string{
		"go mod tidy",
		repotools.InstallPython3(commandUsage),
		"pip3 install virtualenv",
		"python3 -m virtualenv .python-virtualenv",
		"source .python-virtualenv/bin/activate",
		"pip3 install -r requirements.txt",
		"source ./.python-virtualenv/bin/activate",
		repotools.InstallNodeJs(commandUsage),
		"npm install",
	}

	for _, command := range commands {
		ansi.Blue().Printf("Running %s", command).LF().Reset()
		if out, err := runcommand.ShellExecute(command); err != nil {
			ansi.
				Red().
				Printf("Error initializing repo (%s): %v", command, err).
				LF().
				Reset().
				Fatal(exit.GeneralError)
		} else {
			if strings.TrimSpace(out) == "" {
				out = "<no output>"
			}
			ansi.
				Green().
				Printf("Success: '%s'", out).
				LF().
				Reset()
		}
	}
	ansi.Green().Println("init-repo: success").Reset()
}
