package os

import (
	"os"
)

// Executable - Abstract os.Executable
var Executable = os.Executable

// ResetOsExecutableWrapper - Reset our os.Executable wrapper to its original native state
func ResetOsExecutableWrapper() {
	Executable = os.Executable
}
