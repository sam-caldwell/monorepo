package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
)

// PrinterFunction - a simple function pattern for our Printers
// type PrinterFunction func(indexWidth, pos, width int, name, key string)
type PrinterFunction func(format string, args ...any)

//indexWidth, pos, width int, name, key string)

// SelectPrinter - select a color or non-color printer
func SelectPrinter() (PrinterFunction, bool) {
	for _, arg := range os.Args {
		if (arg == "-color") || (arg == "--color") {
			return func(format string, args ...any) {
				ansi.Blue().
					Space().
					Printf(format, args...).
					Reset()
			}, true
		}
	}
	return func(format string, args ...any) {
		fmt.Printf(format, args...)
	}, false
}
