package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	exit2 "github.com/sam-caldwell/monorepo/go/exit"
	"os"
	"strings"
)

func main() {
	exit2.IfVersionRequested()
	if len(os.Args) < 2 {
		ansi.Red().Println("Missing message string").Fatal(exit2.MissingArg)
	}
	morseCode, err := convert.ToMorseCode(strings.Join(os.Args[1:], " "))
	if err != nil {
		ansi.Red().Printf("Error: %s", err).LF().Reset()
	}
	println(morseCode)
}
