package os

import (
	"os"
)

// Chdir - Abstract os.Chdir
var Chdir = ResetOsChdirWrapper()

// ResetOsChdirWrapper - Reset our os.Chdir wrapper to its original native state
func ResetOsChdirWrapper() func(dir string) error {
	Chdir = os.Chdir
	return Chdir
}
