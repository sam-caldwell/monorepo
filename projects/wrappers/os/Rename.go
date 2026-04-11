package os

import (
	"os"
)

// Rename - Abstract os.Rename
var Rename = os.Rename

// ResetOsRenameWrapper - Reset our os.Rename wrapper to its original native state
func ResetOsRenameWrapper() {
	Rename = os.Rename
}
