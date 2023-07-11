package os

import (
	"os"
)

// Stat - Abstract os.Stat
var Stat = os.Stat

// ResetOsStatWrapper - Reset our os.Stat wrapper to its original native state
func ResetOsStatWrapper() {
	Stat = os.Stat
}
