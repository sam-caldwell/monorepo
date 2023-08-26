package os

import (
	"os"
)

// Lstat - Abstract os.Lstat
var Lstat = os.Lstat

// ResetOsLstatWrapper - Reset our os.Lstat wrapper to its original native state
func ResetOsLstatWrapper() {
	Lstat = os.Lstat
}
