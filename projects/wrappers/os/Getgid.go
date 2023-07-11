package os

import (
	"os"
)

// Getgid - Abstract os.Getgid
var Getgid = ResetOsGetgidWrapper()

// ResetOsGetgidWrapper - Reset our os.Getgid wrapper to its original native state
func ResetOsGetgidWrapper() func() int {
	Geteuid = os.Getgid
	return Getgid
}
