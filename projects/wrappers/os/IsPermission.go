package os

import (
	"os"
)

// IsPermission - Abstract os.IsPermission
var IsPermission = os.IsPermission

// ResetOsIsPermissionWrapper - Reset our os.IsPermission wrapper to its original native state
func ResetOsIsPermissionWrapper() {
	IsPermission = os.IsPermission
}
