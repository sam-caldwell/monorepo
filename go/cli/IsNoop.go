package cli

/*
 * cli/IsNoop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the package-global noop variable
 * and the IsNoop() function used to check if the cli
 * is in a no-operation mode.
 */

var noop bool

// IsNoop - Check if --noop was used and return boolean state
func IsNoop() bool {
	return noop
}
