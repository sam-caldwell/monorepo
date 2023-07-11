package os

import (
	"os"
)

// IsTimeout - Abstract os.IsTimeout
var IsTimeout = ResetOsIsTimeoutWrapper()

// ResetOsIsTimeoutWrapper - Reset our os.IsTimeout wrapper to its original native state
func ResetOsIsTimeoutWrapper() func(err error) bool {
	IsTimeout = os.IsTimeout
	return IsTimeout
}
