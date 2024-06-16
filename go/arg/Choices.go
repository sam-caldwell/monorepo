package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Choices - Get a string from command-line and verify it is one of a set of choices
//
//	(c) 2023 Sam Caldwell.  MIT License
func Choices(name, defaultTemplate, usage string, choices ...string) (*string, error) {
	value := flag.String(name, defaultTemplate, usage)
	for _, validValue := range choices {
		if *value != validValue {
			return nil, fmt.Errorf(errors.InvalidInput+errors.Details, name)
		}
	}
	return value, nil
}
