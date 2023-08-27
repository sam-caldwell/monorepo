package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/ansi"
)

// SkipMessagePrintFunc - the signature for our skip printer function
type SkipMessagePrintFunc func(name, subject, msg string)

// SkipMessagePrinter - Show a "skip" message
func SkipMessagePrinter(programName string, useColor, quietMode bool, counter *int) SkipMessagePrintFunc {
	return func(group, test, message string) {
		*counter++
		const format = "%s (%s) [SKIP](%s): %s\n"
		if !quietMode {
			if useColor {
				ansi.Yellow().Printf(format, programName, group, test, message).Reset()
			} else {
				fmt.Printf(format+"", programName, group, test, message)
			}
		}
	}
}
