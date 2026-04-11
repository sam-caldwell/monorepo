package os

import (
	"os"
)

// Symlink - Abstract os.Symlink
var Symlink = os.Symlink

// ResetOsSymlinkWrapper - Reset our os.Symlink wrapper to its original native state
func ResetOsSymlinkWrapper() {
	Symlink = os.Symlink
}
