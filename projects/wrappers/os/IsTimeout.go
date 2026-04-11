package os

import (
	"os"
)

// IsTimeout - Abstract os.IsTimeout
var IsTimeout = os.IsTimeout

// ResetOsIsTimeoutWrapper - Reset our os.IsTimeout wrapper to its original native state
func ResetOsIsTimeoutWrapper() {
	IsTimeout = os.IsTimeout
}
