package os

import (
	"os"
)

// SameFile - Abstract os.SameFile
var SameFile = os.SameFile

// ResetOsSameFileWrapper - Reset our os.SameFile wrapper to its original native state
func ResetOsSameFileWrapper() {
	SameFile = os.SameFile
}
