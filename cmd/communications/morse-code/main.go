package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"os"
	"strings"
)

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 2 {
		ansi.Red().Println("Missing message string").Fatal(exit.MissingArg)
	}
	morseCode, err := convert.ToMorseCode(strings.Join(os.Args[1:], " "))
	if err != nil {
		ansi.Red().Printf("Error: %s", err).LF().Reset()
	}
	println(morseCode)
}
