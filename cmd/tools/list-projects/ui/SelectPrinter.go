package ui

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
)

// SelectPrinter - select a color or non-color printer
func SelectPrinter() (PrinterFunction, bool) {
	for _, arg := range os.Args {
		if arg == "-color" {
			return func(width int, name, key string) {
				ansi.Blue().
					Printf("%*s", width, name).
					Space().
					Yellow().
					Printf("%s", key).LF().
					Reset()
			}, true
		}
	}
	return func(width int, name, key string) {
		fmt.Printf("%*s %s\n", width, name, key)
	}, false
}
