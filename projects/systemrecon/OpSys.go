package systemrecon

/*
 * OpSys ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * OpSys()
 *
 * 	Return the current operating system (using GOOS)
 */

import "runtime"

// OpSys - Return the operating system (e.g. windows, darwin, linux)
func OpSys() (string, error) {

	return runtime.GOOS, nil

}
