package systemrecon

import "runtime"

// OpSys - Return the operating system (e.g. windows, darwin, linux)
func OpSys() (string, error) {
	return runtime.GOOS, nil
}
