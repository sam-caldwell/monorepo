package exit

import (
	"fmt"
	"os"
)

// OnCondition - A standard way to quickly terminate a program.
func OnCondition(condition bool, code int, err string, usage string) {
	if condition {
		var usageMsg string
		if usage != "" {
			usageMsg = fmt.Sprintf(errUsage, usage)
		}
		fmt.Printf(errMessage, err, usageMsg)
		os.Exit(code)
	}
}
