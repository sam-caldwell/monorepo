package os

import (
	"os"
)

// Chmod - Abstract os.Chmod
var Chmod = ResetOsChmodWrapper()

// ResetOsChmodWrapper - Reset our os.Chmod wrapper to its original native state
func ResetOsChmodWrapper() func(name string, mode os.FileMode) error {
	Chmod = os.Chmod
	return Chmod
}
