package os

import (
	"os"
)

// Clearenv - Abstract os.Clearenv
var Clearenv = ResetOsClearenvWrapper()

// ResetOsClearenvWrapper - Reset our os.Clearenv wrapper to its original native state
func ResetOsClearenvWrapper() func() {
	Clearenv = os.Clearenv
	return Clearenv
}
