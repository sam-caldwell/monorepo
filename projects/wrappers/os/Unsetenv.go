package os

import (
	"os"
)

// Unsetenv - Abstract os.Unsetenv
var Unsetenv = os.Unsetenv

// ResetOsUnsetenvWrapper - Reset our os.Unsetenv wrapper to its original native state
func ResetOsUnsetenvWrapper() {
	Unsetenv = os.Unsetenv
}
