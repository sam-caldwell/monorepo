package monorepo

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/version"
)

// GetArgs - Process commandline arguments
func GetArgs() (commands []string, debug *bool) {

	cmdVersion := flag.Bool("version", false, "show version")

	debug = flag.Bool("debug", false, "print debug messages")

	flag.Parse()

	commands = flag.Args()

	if *cmdVersion {
		version.Show()
	}

	if *debug {
		ansi.Blue().
			Println("Debug: enabled").
			Printf("\tcommand: %v\n", commands).
			LF().
			Reset()
	}
	return
}
