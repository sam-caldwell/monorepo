package os

import (
	"os"
)

// UserCacheDir - Abstract os.UserCacheDir
var UserCacheDir = os.UserCacheDir

// ResetOsUserCacheDirWrapper - Reset our os.UserCacheDir wrapper to its original native state
func ResetOsUserCacheDirWrapper() {
	UserCacheDir = os.UserCacheDir
}
