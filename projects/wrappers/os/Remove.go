package os

import (
	"os"
)

// Remove - Abstract os.Remove
var Remove = os.Remove

// ResetOsRemoveWrapper - Reset our os.Remove wrapper to its original native state
func ResetOsRemoveWrapper() {
	Remove = os.Remove
}
