package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc"
)

// ShowStats - show stats in UI
func ShowStats(programName string, width int, useColor, quietMode bool, countPass, countFail, countSkip int) {
	if quietMode {
		return
	}
	misc.ShowStats(width, useColor,
		"Linter Stats",
		fmt.Sprintf("  Total:%6d", countPass+countSkip+countFail),
		map[string]int{
			"pass": countPass,
			"fail": countFail,
			"skip": countSkip,
		})
}
