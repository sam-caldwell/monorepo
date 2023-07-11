package os

import (
	"os"
)

// IsPathSeparator - Abstract os.IsPathSeparator
var IsPathSeparator = ResetOsIsPathSeparatorWrapper()

// ResetOsIsPathSeparatorWrapper - Reset our os.IsPathSeparator wrapper to its original native state
func ResetOsIsPathSeparatorWrapper() func(c uint8) bool {
	IsPathSeparator = os.IsPathSeparator
	return IsPathSeparator
}
