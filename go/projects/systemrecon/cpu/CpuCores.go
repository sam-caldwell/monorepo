package systemrecon

/*
 * CpuCores
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file returns the number of CPU cores
 */

import (
	"runtime"
)

// CpuCores - Return the CPU core count
func CpuCores() (int, error) {
	return runtime.NumCPU(), nil
}
