package os

import (
	"os"
)

// WriteFile - Abstract os.WriteFile
var WriteFile = os.WriteFile

// ResetOsWriteFileWrapper - Reset our os.WriteFile wrapper to its original native state
func ResetOsWriteFileWrapper() {
	WriteFile = os.WriteFile
}
