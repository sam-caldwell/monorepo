package os

import (
	"os"
)

// ReadDir - Abstract os.ReadDir
var ReadDir = os.ReadDir

// ResetOsReadDirWrapper - Reset our os.ReadDir wrapper to its original native state
func ResetOsReadDirWrapper() {
	ReadDir = os.ReadDir
}
