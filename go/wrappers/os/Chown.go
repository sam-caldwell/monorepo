package os

import (
	"os"
)

// Chown - Abstract os.Chown
var Chown = os.Chown

// ResetOsChownWrapper - Reset our os.Chown wrapper to its original native state
func ResetOsChownWrapper() {
	Chown = os.Chown
}
