package os

import (
	"os"
)

// Chdir - Abstract os.Chdir
var Chdir = os.Chdir

// ResetOsChdirWrapper - Reset our os.Chdir wrapper to its original native state
func ResetOsChdirWrapper() {
	Chdir = os.Chdir
}
