package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	terminal "github.com/sam-caldwell/monorepo/go/terminal/widgets"
	"strings"
)

func showComment(comment string, nestLevel int) {
	vChar := ""
	if nestLevel != 0 {
		vChar = "│"
	}
	if strings.TrimSpace(comment) != "" {
		ansi.Cyan().Printf("      %s └┬─ /*\n", vChar)
		for _, rawLine := range strings.Split(comment, "\n") {
			for _, line := range terminal.SplitStringIntoLines(rawLine, 80) {
				ansi.Printf("      %s  │   * ", vChar).Yellow().Printf("%s\n", line).Cyan()
			}
		}
		ansi.Cyan().Printf("      %s  └─  */\n", vChar)
	}
}
