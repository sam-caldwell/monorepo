package os

import (
	"os"
)

// UserConfigDir - Abstract os.UserConfigDir
var UserConfigDir = os.UserConfigDir

// ResetOsUserConfigDirWrapper - Reset our os.UserConfigDir wrapper to its original native state
func ResetOsUserConfigDirWrapper() {
	UserConfigDir = os.UserConfigDir
}
