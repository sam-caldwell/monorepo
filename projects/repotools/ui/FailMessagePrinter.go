package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
)

// FailMessagePrintFunc - the signature for our Fail printer function
type FailMessagePrintFunc func(group, test string, err error)

// FailMessagePrinter - Show a "Fail" message
func FailMessagePrinter(programName string, useColor bool, counter *int) FailMessagePrintFunc {
	return func(group, test string, err error) {
		if err == nil {
			return
		}
		*counter++
		const format = "[FAIL](%s)(%s)(%s) %s"
		if useColor {
			ansi.Red().Printf(format, programName, group, test, err).LF().Reset()
		} else {
			fmt.Printf(format, programName, group, test, err)
		}
	}
}