package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
)

// NoticeMessagePrintFunc - the signature for our notice printer function
type NoticeMessagePrintFunc func(format string, args ...any)

// NoticeMessagePrinter - show a notice message
func NoticeMessagePrinter(programName string, useColor, quietMode bool) NoticeMessagePrintFunc {
	return func(format string, args ...any) {
		if quietMode {
			return
		}
		if useColor {
			ansi.Yellow().Printf(">>"+format, args...).LF().Reset()
		} else {
			fmt.Printf(format+"\n", args...)
		}
	}
}
