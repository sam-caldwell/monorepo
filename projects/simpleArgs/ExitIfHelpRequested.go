package simpleArgs

import (
	"github.com/sam-caldwell/go/v2/projects/exit"
	"os"
)

// ExitIfHelpRequested - intercept a request for help information
func ExitIfHelpRequested(commandUsage string) {
	for _, arg := range os.Args {
		exit.OnCondition(arg == "-h" || arg == "--help", 0, "", commandUsage)
	}
}
