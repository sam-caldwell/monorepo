package arg

import "flag"

// Debug - Return true if -debug flag is passed to command-line
//
//	(c) 2023 Sam Caldwell.  MIT License
func Debug() *bool {
	value := flag.Bool("debug", false, "set the debug flag")
	return value
}
