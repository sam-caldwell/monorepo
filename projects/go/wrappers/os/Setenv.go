package os

import (
	"os"
)

// Setenv - Abstract os.Setenv
var Setenv = os.Setenv

// ResetOsSetenvWrapper - Reset our os.Setenv wrapper to its original native state
func ResetOsSetenvWrapper() {
	Setenv = os.Setenv
}
