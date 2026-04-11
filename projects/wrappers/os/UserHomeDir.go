package os

import (
	"os"
)

// UserHomeDir - Abstract os.UserHomeDir
var UserHomeDir = os.UserHomeDir

// ResetOsUserHomeDirWrapper - Reset our os.UserHomeDir wrapper to its original native state
func ResetOsUserHomeDirWrapper() {
	UserHomeDir = os.UserHomeDir
}
