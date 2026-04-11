package os

import (
	"os"
)

// NewFile - Abstract os.NewFile
var NewFile = os.NewFile

// ResetOsNewFileWrapper - Reset our os.NewFile wrapper to its original native state
func ResetOsNewFileWrapper() {
	NewFile = os.NewFile
}
