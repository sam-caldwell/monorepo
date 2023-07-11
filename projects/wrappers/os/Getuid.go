package os

import (
	"os"
)

// Getuid - Abstract os.Getuid
var Getuid = ResetOsGetuidWrapper()

// ResetOsGetuidWrapper - Reset our os.Getuid wrapper to its original native state
func ResetOsGetuidWrapper() func() int {
	Getuid = os.Getuid
	return Getuid
}
