package os

import (
	"os"
)

// TempDir - Abstract os.TempDir
var TempDir = os.TempDir

// ResetOsTempDirWrapper - Reset our os.TempDir wrapper to its original native state
func ResetOsTempDirWrapper() {
	TempDir = os.TempDir
}
