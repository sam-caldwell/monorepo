package os

import (
	"os"
)

// Chmod - Abstract os.Chmod
var Chmod = os.Chmod

// ResetOsChmodWrapper - Reset our os.Chmod wrapper to its original native state
func ResetOsChmodWrapper() {
	Chmod = os.Chmod
}
