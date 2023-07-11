package os

import (
	"os"
)

// IsPathSeparator - Abstract os.IsPathSeparator
var IsPathSeparator = os.IsPathSeparator

// ResetOsIsPathSeparatorWrapper - Reset our os.IsPathSeparator wrapper to its original native state
func ResetOsIsPathSeparatorWrapper() {
	IsPathSeparator = os.IsPathSeparator
}
