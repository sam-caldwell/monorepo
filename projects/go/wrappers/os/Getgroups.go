package os

import (
	"os"
)

// Getgroups - Abstract os.Getgroups
var Getgroups = os.Getgroups

// ResetOsGetgroupsWrapper - Reset our os.Getgroups wrapper to its original native state
func ResetOsGetgroupsWrapper() {
	Getgroups = os.Getgroups
}
