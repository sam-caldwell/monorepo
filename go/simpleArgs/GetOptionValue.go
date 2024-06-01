package simpleArgs

import (
	"fmt"
	cliArgs "github.com/sam-caldwell/monorepo/go/misc/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"strings"
)

// GetOptionValue returns the value (next arg) if the named option exists.
func GetOptionValue(name string) (value string, err error) {

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {

		option := args[i]

		if option == name {

			if i+1 >= len(args) || strings.HasPrefix(args[i+1], cliArgs.DoubleHyphen) {

				return words.EmptyString, fmt.Errorf(OptionHasNoValue)

			}

			return args[i+1], nil

		}

	}

	return "", fmt.Errorf(OptionNotFound)

}
