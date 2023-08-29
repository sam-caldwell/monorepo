package os

import (
	"os"
)

// Environ - Abstract os.Environ
var Environ = os.Environ

// ResetOsEnvironWrapper - Reset our os.Environ wrapper to its original native state
func ResetOsEnvironWrapper() {
	Environ = os.Environ
}
