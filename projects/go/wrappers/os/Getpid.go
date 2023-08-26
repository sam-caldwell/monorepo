package os

import (
	"os"
)

// Getpid - Abstract os.Getpid
var Getpid = os.Getpid

// ResetOsGetpidWrapper - Reset our os.Getpid wrapper to its original native state
func ResetOsGetpidWrapper() {
	Getpid = os.Getpid
}
