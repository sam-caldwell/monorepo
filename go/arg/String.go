package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"regexp"
)

// String - get a string from the command-line and validate it against a regular expression pattern
//
//	(c) 2023 Sam Caldwell.  MIT License
func String(name, defaultValue, usage, validationString string) (*string, error) {
	pattern := regexp.MustCompile(validationString)
	value := flag.String(name, defaultValue, usage)
	if pattern.MatchString(*value) {
		return value, nil
	}
	return nil, fmt.Errorf(errors.InvalidInput+errors.Details, name)
}
