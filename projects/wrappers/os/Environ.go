package os

import (
	"os"
)

// Environ - Abstract os.Environ
var Environ = ResetOsEnvironWrapper()

// ResetOsEnvironWrapper - Reset our os.Environ wrapper to its original native state
func ResetOsEnvironWrapper() func() []string {
	Environ = os.Environ
	return Environ
}
