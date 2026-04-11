package os

import (
	"os"
)

// Getenv - Abstract os.Getenv
var Getenv = os.Getenv

// ResetOsGetenvWrapper - Reset our os.Getenv wrapper to its original native state
func ResetOsGetenvWrapper() {
	Getenv = os.Getenv
}
