package os

import (
	"os"
)

// Executable - Abstract os.Executable
var Executable = ResetOsExecutableWrapper()

// ResetOsExecutableWrapper - Reset our os.Executable wrapper to its original native state
func ResetOsExecutableWrapper() func() (string, error) {
	Executable = os.Executable
	return Executable
}
