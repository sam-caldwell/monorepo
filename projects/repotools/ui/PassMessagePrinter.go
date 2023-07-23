package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
)

// PassMessagePrintFunc - the signature for our Pass printer function
type PassMessagePrintFunc func(subject, message string)

// PassMessagePrinter - Show a "Pass" message
func PassMessagePrinter(programName string, useColor bool, counter *int) PassMessagePrintFunc {
	return func(subject, message string) {
		*counter++
		const format = "%s (%s) [PASS](%s)"
		if useColor {
			ansi.Green().Printf(format, programName, subject, message).LF().Reset()
		} else {
			fmt.Printf(format, programName, subject, message)
		}
	}
}
