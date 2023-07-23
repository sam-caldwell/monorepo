package repocli

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc"
)

func ShowStats(programName string, width int, useColor bool, countPass, countFail, countSkip int) {
	misc.ShowStats(width, useColor,
		"Linter Stats",
		fmt.Sprintf("  Total:%6d", countPass+countSkip+countFail),
		map[string]int{
			"pass": countPass,
			"fail": countFail,
			"skip": countSkip,
		})
}
