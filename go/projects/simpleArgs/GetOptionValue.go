package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
	"os"
	"strings"
)

// GetOptionValue returns the value (next arg) if the named option exists.
func GetOptionValue(name string) (value string, err error) {
	const (
		doubleHyphen     = "--"
		optionHasNoValue = "option has no value"
		optionNotFound   = "option not found"
	)

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {

		option := args[i]

		if option == name {

			if i+1 >= len(args) || strings.HasPrefix(args[i+1], doubleHyphen) {

				return words.EmptyString, fmt.Errorf(optionHasNoValue)

			}

			return args[i+1], nil

		}

	}

	return "", fmt.Errorf(optionNotFound)

}
