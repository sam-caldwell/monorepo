package os

import (
	"os"
)

// Chown - Abstract os.Chown
var Chown = ResetOsChownWrapper()

// ResetOsChownWrapper - Reset our os.Chown wrapper to its original native state
func ResetOsChownWrapper() func(name string, uid, gid int) error {
	Chown = os.Chown
	return Chown
}
