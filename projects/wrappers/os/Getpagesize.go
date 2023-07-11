package os

import (
	"os"
)

// Getpagesize - Abstract os.Getpagesize
var Getpagesize = ResetOsGetpagesizeWrapper()

// ResetOsGetpagesizeWrapper - Reset our os.Getpagesize wrapper to its original native state
func ResetOsGetpagesizeWrapper() func() int {
	Getpagesize = os.Getpagesize
	return Getpagesize
}
