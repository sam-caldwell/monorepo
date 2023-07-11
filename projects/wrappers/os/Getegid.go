package os

import (
	"os"
)

// Getegid - Abstract os.Getegid
var Getegid = os.Getegid

// ResetOsGetegidWrapper - Reset our os.Getegid wrapper to its original native state
func ResetOsGetegidWrapper() {
	Getegid = os.Getegid
}
