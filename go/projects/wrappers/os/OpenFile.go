package os

import (
	"os"
)

// OpenFile - Abstract os.OpenFile
var OpenFile = os.OpenFile

// ResetOsOpenFileWrapper - Reset our os.OpenFile wrapper to its original native state
func ResetOsOpenFileWrapper() {
	OpenFile = os.OpenFile
}
