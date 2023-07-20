package ui

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
	"strings"
)

// UseHeadersAndFooters determine if we should pretty print
func UseHeadersAndFooters() bool {
	for _, arg := range os.Args {
		if arg == "-banner" {
			return true
		}
	}
	return false
}

func PrintHeader(msg string, useColor bool, displayWidth int) {
	if useColor {
		ansi.Blue().
			Line("-", displayWidth).
			Space().Println(msg).
			Line("-", displayWidth).
			Reset()
	} else {
		fmt.Println(strings.Repeat("-", displayWidth))
		fmt.Println(msg)
		fmt.Println(strings.Repeat("-", displayWidth))
	}
}

func PrintFooter(msg string, useColor bool, displayWidth int) {
	if useColor {
		ansi.Blue().
			LF().
			Line("-", displayWidth).
			Space().Print(msg).LF().
			Line("-", displayWidth).
			Reset()
	} else {
		fmt.Println(strings.Repeat("-", displayWidth))
		fmt.Println(msg)
		fmt.Println(strings.Repeat("-", displayWidth))

	}
}
