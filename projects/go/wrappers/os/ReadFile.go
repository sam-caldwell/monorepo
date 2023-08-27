package os

import (
	"os"
)

// ReadFile - Abstract os.ReadFile
var ReadFile = os.ReadFile

// ResetOsReadFileWrapper - Reset our os.ReadFile wrapper to its original native state
func ResetOsReadFileWrapper() {
	ReadFile = os.ReadFile
}
