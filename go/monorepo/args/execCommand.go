package args

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// execCommand - Parse the command-line arguments for the exec command
//
//	 Expected syntax:
//
//			monorepo exec <code string> [parameters]
//
//	The input args []string starts with the first arg (exec)
//	and all subsequent arguments.
func (arg *Arguments) execCommand(args []string) {
	exit.OnCondition(len(args) < 3, exit.MissingArg, "Missing arguments for exec command", checkUsage)
	for i, a := range args {
		fmt.Printf("[%d]: %s\n", i, a)
	}

	arg.group = cmdExec // exec (args[1])

	args = arg.generalOptions(args[2:])

	arg.command = args[0] // this is the code string

	fmt.Printf("group: %s\n", arg.group)
	fmt.Printf("command: %s\n", arg.command)
	fmt.Printf("parameters: %v\n", arg.parameters)

	var parameterName string
	var expectValue bool

	for _, parameter := range arg.generalOptions(args[1:]) {
		switch parameter {
		case optionWorkingDir:
			parameterName = strings.TrimLeft(parameter, words.Hyphen)
			expectValue = true
		case optionShell:
			parameterName = strings.TrimLeft(parameter, words.Hyphen)
			expectValue = true
		default:
			if expectValue {
				arg.parameters[parameterName] = parameter
				expectValue = false
			} else {
				arg.syntaxError(fmt.Sprintf("Expected value for %s", parameter))
			}
		}
	}
}
