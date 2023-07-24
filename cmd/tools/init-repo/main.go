package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	"strings"
)

func main() {

	commands := []string{
		"go mod tidy",
		"pip3 install -r requirements.txt",
		"npm install",
	}

	ansi.Blue().Println("init-repo: starting")

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
