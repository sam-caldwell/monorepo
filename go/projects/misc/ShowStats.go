package misc

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi"
	"strings"
)

func ShowStats(width int, useColor bool, header, footer string, data map[string]int) {
	var content string
	keyWidth := 0
	for key := range data {
		if w := len(key); w > keyWidth {
			keyWidth = w
		}
	}
	keyWidth += 2
	for key, value := range data {
		content += fmt.Sprintf("  %*s: %6d\n", keyWidth, key, value)
	}
	singleLine := strings.Repeat("-", width) + "\n"
	doubleLine := strings.Repeat("=", width) + "\n"
	if useColor {
		ansi.Blue().
			Printf("\n"+
				singleLine+
				" %s\n"+
				singleLine+
				"%s"+
				singleLine+
				" %s\n"+
				doubleLine,
				header, content, footer).
			LF().
			Reset()
	} else {
		fmt.Printf("\n"+
			singleLine+
			" %s\n"+
			singleLine+
			"%s"+
			singleLine+
			" %s"+
			doubleLine,
			header, content, footer)
	}
}
