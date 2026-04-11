package arg

import (
	"flag"
)

// Bool - get bool from the command-line and validate it against the given parameters
//
//	(c) 2023 Sam Caldwell.  MIT License
func Bool(name string, defaultValue bool, usage string) (*bool, error) {
	value := flag.Bool(name, defaultValue, usage)
	return value, nil
}
