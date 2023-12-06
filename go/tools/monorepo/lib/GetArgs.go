package monorepo

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/version"
)

// GetArgs - Process commandline arguments
func GetArgs() (command *string, debug *bool) {
	command = GetCommand()

	cmdVersion := flag.Bool("version", false, "show version")

	debug = flag.Bool("debug", false, "print debug messages")

	flag.Parse()
	if *cmdVersion {
		version.Show()
	}

	if *debug {
		ansi.Blue().
			Println("Debug:").
			Printf("\tcommand: %s\n", *command).
			LF().
			Reset()
	}
	return
}
