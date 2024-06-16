package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Uint64 - get uint from the command-line and validate it against the given parameters
//
//	(c) 2023 Sam Caldwell.  MIT License
func Uint64(name string, defaultValue, min, max uint64, usage string) (*uint64, error) {
	value := flag.Uint64(name, defaultValue, usage)
	if *value <= min || *value >= max {
		return nil, fmt.Errorf(errors.InvalidInput+errors.Details, name)
	}
	return value, nil
}
