package os

import (
	"os"
)

// Getpid - Abstract os.Getpid
var Getpid = ResetOsGetpidWrapper()

// ResetOsGetpidWrapper - Reset our os.Getpid wrapper to its original native state
func ResetOsGetpidWrapper() func() int {
	Getpid = os.Getpid
	return Getpid
}
