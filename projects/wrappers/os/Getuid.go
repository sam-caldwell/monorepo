package os

import (
	"os"
)

// Getuid - Abstract os.Getuid
var Getuid = os.Getuid

// ResetOsGetuidWrapper - Reset our os.Getuid wrapper to its original native state
func ResetOsGetuidWrapper() {
	Getuid = os.Getuid
}
