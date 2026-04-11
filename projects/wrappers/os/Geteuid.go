package os

import (
	"os"
)

// Geteuid - Abstract os.Geteuid
var Geteuid = os.Geteuid

// ResetOsGeteuidWrapper - Reset our os.Geteuid wrapper to its original native state
func ResetOsGeteuidWrapper() {
	Geteuid = os.Geteuid
}
