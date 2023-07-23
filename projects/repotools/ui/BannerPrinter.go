package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
)

// BannerPrinterFunc - the signature for our banner printer function
type BannerPrinterFunc func(color *ansi.Color, args ...any)

// BannerMessagePrinter - show a banner message
func BannerMessagePrinter(programName string, useColor, quietMode bool, width int) BannerPrinterFunc {
	return func(color *ansi.Color, args ...any) {
		if quietMode {
			return
		}
		if useColor {
			color.
				Line("-", width)
			for _, line := range args {
				ansi.Printf("  %v", line).LF()
			}
			ansi.
				Line("-", width).
				Reset()
		} else {
			for _, line := range args {
				fmt.Printf("%v\n", line)
			}
		}
	}
}
