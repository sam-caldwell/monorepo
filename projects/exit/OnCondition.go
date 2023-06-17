package exit

import (
	"fmt"
	"os"
)

// OnCondition - A standard way to quickly terminate a program.
func OnCondition(condition bool, code int, err string, usage string) {
	if condition {
		fmt.Printf(errMessage, err, fmt.Sprintf(errUsage, usage))
		os.Exit(code)
	}
}
