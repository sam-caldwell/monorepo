package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/list"
	cliArgs "github.com/sam-caldwell/monorepo/go/misc/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"strings"
)

// GetCommand - Given commandline args (os.Args) return the parsed command or handle --help and --version
//
//	(c) 2023 Sam Caldwell.  MIT License
func GetCommand(helpText string) (command string) {
	exit.OnCondition(len(os.Args) < 2, exit.InvalidInput, errors.MissingArguments, helpText)
	command = strings.ToLower(os.Args[1])
	exit.OnCondition((command == cliArgs.Help) || (command == cliArgs.H), 0, words.Help, helpText)
	if command == cliArgs.Version {
		fmt.Println(version.Version)
		os.Exit(exit.Success)
	}
	list.DeleteElement(os.Args, exit.GeneralError)
	return command
}
