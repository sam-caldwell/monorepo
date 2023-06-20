package systemrecon

import "runtime"

// CpuArch - Return CPU Architecture (e.g. amd64)
func CpuArch() (string, error) {
	return runtime.GOARCH, nil
}
