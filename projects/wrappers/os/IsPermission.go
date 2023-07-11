package os

import (
	"os"
)

// IsPermission - Abstract os.IsPermission
var IsPermission = ResetOsIsPermissionWrapper()

// ResetOsIsPermissionWrapper - Reset our os.IsPermission wrapper to its original native state
func ResetOsIsPermissionWrapper() func(err error) bool {
	IsPermission = os.IsPermission
	return IsPermission
}
