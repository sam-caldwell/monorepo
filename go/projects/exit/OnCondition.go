package exit

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
)

// OnCondition - A standard way to quickly terminate a program.
func OnCondition(condition bool, code int, err string, usage string) {
	if condition {
		var usageMsg string
		var formatString string
		if usage != "" {
			usageMsg = fmt.Sprintf(errUsage, usage)
		}
		if err == "" {
			formatString = "%s%s\n"
		} else {
			formatString = errMessage
		}
		fmt.Printf(formatString, err, usageMsg)
		os.Exit(code)
	}
}
