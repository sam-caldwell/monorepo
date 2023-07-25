package os

import (
	"os"
)

// Readlink - Abstract os.Readlink
var Readlink = os.Readlink

// ResetOsReadlinkWrapper - Reset our os.Readlink wrapper to its original native state
func ResetOsReadlinkWrapper() {
	Readlink = os.Readlink
}
