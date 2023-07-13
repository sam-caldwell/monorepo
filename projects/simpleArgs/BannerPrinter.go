package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
	"strings"
)

// BannerPrinter - If os.Args contains --banner or -banner, return a banner printer function
func BannerPrinter() func(bannerText string, displayWidth int, margin int, useColor bool) {
	var showBanner = false
	for _, arg := range os.Args {
		if (arg == "-banner") || (arg == "--banner") {
			showBanner = true
		}
	}
	if showBanner {
		return func(bannerText string, displayWidth int, margin int, useColor bool) {
			if !showBanner {
				return //nop
			}

			if useColor {
				ansi.Green().
					Line("-", displayWidth).
					Space().Print(bannerText).LF().
					Line("-", displayWidth).
					Reset()
			} else {
				fmt.Println(strings.Repeat("-", displayWidth))
				fmt.Println(bannerText)
				fmt.Println(strings.Repeat("-", displayWidth))
			}
		}
	}
	return func(bannerText string, displayWidth int, margin int, useColor bool) {}
}
