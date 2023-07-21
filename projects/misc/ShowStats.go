package misc

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
)

func ShowStats(useColor bool, header, footer string, data map[string]int) {
	var content string
	keyWidth := 0
	for key, _ := range data {
		if w := len(key); w > keyWidth {
			keyWidth = w
		}
	}
	keyWidth += 2
	for key, value := range data {
		content += fmt.Sprintf("  %*s: %6d\n", keyWidth, key, value)
	}
	if useColor {
		ansi.Blue().
			Printf("\n"+
				"------------------\n"+
				" %s\n"+
				"------------------\n"+
				"%s"+
				"------------------\n"+
				" %s\n"+
				"==================\n",
				header, content, footer).
			LF().
			Reset()
	} else {
		fmt.Printf("\n"+
			"------------------\n"+
			" %s\n"+
			"------------------\n"+
			"%s"+
			"------------------\n"+
			" %s"+
			"==================\n",
			header, content, footer)
	}
}
