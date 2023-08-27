package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/ansi"
)

// PassMessagePrintFunc - the signature for our Pass printer function
type PassMessagePrintFunc func(subject, message string)

// PassMessagePrinter - Show a "Pass" message
func PassMessagePrinter(programName string, useColor bool, counter *int) PassMessagePrintFunc {
	return func(subject, message string) {
		*counter++
		const format = "[PASS](%s)(%s) %s\n"
		if useColor {
			ansi.Green().Printf(format, programName, subject, message).Reset()
		} else {
			fmt.Printf(format, programName, subject, message)
		}
	}
}
