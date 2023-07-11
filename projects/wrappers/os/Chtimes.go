package os

import (
	"os"
	"time"
)

// Chtimes - Abstract os.Chtimes
var Chtimes = ResetOsChtimesWrapper()

// ResetOsChtimesWrapper - Reset our os.Chtimes wrapper to its original native state
func ResetOsChtimesWrapper() func(name string, atime time.Time, mtime time.Time) error {
	Chtimes = os.Chtimes
	return Chtimes
}
