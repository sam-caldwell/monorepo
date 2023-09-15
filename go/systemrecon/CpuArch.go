package systemrecon

/*
 * CpuArch
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file returns the CPU Architecture (e.g. AMD64, ARM64)
 */

import "runtime"

// CpuArch - Return CPU Architecture (e.g. amd64)
func CpuArch() (string, error) {
	return runtime.GOARCH, nil
}
