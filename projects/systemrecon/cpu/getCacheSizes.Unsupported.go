//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

import (
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(executor runcommand.CommandExecutor, level int) (size int, err error) {
	return 0, nil
}
