package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

// GetCommand - Given commandline args (os.Args) return the parsed command or handle --help and --version
func GetCommand(helpText string) (command string) {
	exit.OnCondition(len(os.Args) < 2, exit.InvalidInput, errors.MissingArguments, helpText)
	command = os.Args[1]
	exit.OnCondition((command == "--help") || (command == "-h"), 0, "Help", helpText)
	if command == "--version" {
		fmt.Println(version.Version)
		os.Exit(0)
	}
	return command
}
