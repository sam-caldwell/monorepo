package os

import (
	"os"
)

// NewFile - Abstract os.NewFile
var NewFile = ResetOsNewFileWrapper()

// ResetOsNewFileWrapper - Reset our os.NewFile wrapper to its original native state
func ResetOsNewFileWrapper() func(fd uintptr, name string) *os.File {
	NewFile = os.NewFile
	return NewFile
}
