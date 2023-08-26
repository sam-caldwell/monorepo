package repocli

import (
	"fmt"
	ansi2 "github.com/sam-caldwell/go/v2/projects/go/ansi"
)

// BannerPrinterFunc - the signature for our banner printer function
type BannerPrinterFunc func(color *ansi2.Color, args ...any)

// BannerMessagePrinter - show a banner message
func BannerMessagePrinter(programName string, useColor, quietMode bool, width int) BannerPrinterFunc {
	return func(color *ansi2.Color, args ...any) {
		if quietMode {
			ansi2.Reset()
			return
		}
		if useColor {
			color.
				Line("-", width)
			for _, line := range args {
				ansi2.Printf("  %v", line).LF()
			}
			ansi2.Line("-", width).
				Reset()
		} else {
			for _, line := range args {
				fmt.Printf("%v\n", line)
			}
		}
		ansi2.White().Reset()
	}
}
