package os

import (
	"os"
)

// Getgroups - Abstract os.Getgroups
var Getgroups = ResetOsGetgroupsWrapper()

// ResetOsGetgroupsWrapper - Reset our os.Getgroups wrapper to its original native state
func ResetOsGetgroupsWrapper() func() ([]int, error) {
	Getgroups = os.Getgroups
	return Getgroups
}
