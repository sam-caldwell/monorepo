package os

import (
	"os"
)

// Getegid - Abstract os.Getegid
var Getegid = ResetOsGetegidWrapper()

// ResetOsGetegidWrapper - Reset our os.Getegid wrapper to its original native state
func ResetOsGetegidWrapper() func() int {
	Getegid = os.Getegid
	return Getegid
}
