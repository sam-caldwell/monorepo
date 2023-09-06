package os

import (
	"os"
)

// Getpagesize - Abstract os.Getpagesize
var Getpagesize = os.Getpagesize

// ResetOsGetpagesizeWrapper - Reset our os.Getpagesize wrapper to its original native state
func ResetOsGetpagesizeWrapper() {
	Getpagesize = os.Getpagesize
}
