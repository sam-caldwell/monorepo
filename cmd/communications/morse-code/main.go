package main

import (
	"github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/convert"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
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
