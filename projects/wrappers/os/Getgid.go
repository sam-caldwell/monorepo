package os

import (
	"os"
)

// Getgid - Abstract os.Getgid
var Getgid = os.Getgid

// ResetOsGetgidWrapper - Reset our os.Getgid wrapper to its original native state
func ResetOsGetgidWrapper() {
	Geteuid = os.Getgid
}
