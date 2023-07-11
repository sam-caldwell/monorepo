package os

import (
	"os"
)

// Geteuid - Abstract os.Geteuid
var Geteuid = ResetOsGeteuidWrapper()

// ResetOsGeteuidWrapper - Reset our os.Geteuid wrapper to its original native state
func ResetOsGeteuidWrapper() func() int {
	Geteuid = os.Geteuid
	return Geteuid
}
