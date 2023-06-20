package systemrecon

import (
	"runtime"
)

// CpuCores - Return the CPU core count
func CpuCores() (int, error) {
	return runtime.NumCPU(), nil
}
