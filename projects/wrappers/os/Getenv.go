package os

import (
	"os"
)

// Getenv - Abstract os.Getenv
var Getenv = ResetOsGetenvWrapper()

// ResetOsGetenvWrapper - Reset our os.Getenv wrapper to its original native state
func ResetOsGetenvWrapper() func(key string) string {
	Getenv = os.Getenv
	return Getenv
}
