package os

import (
	"os"
)

// OpenFile - Abstract os.OpenFile
var OpenFile = ResetOsOpenFileWrapper()

// ResetOsOpenFileWrapper - Reset our os.OpenFile wrapper to its original native state
func ResetOsOpenFileWrapper() func(name string, flag int, perm os.FileMode) (*os.File, error) {
	OpenFile = os.OpenFile
	return OpenFile
}
