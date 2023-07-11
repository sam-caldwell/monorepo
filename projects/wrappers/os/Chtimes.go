package os

import (
	"os"
)

// Chtimes - Abstract os.Chtimes
var Chtimes = os.Chtimes

// ResetOsChtimesWrapper - Reset our os.Chtimes wrapper to its original native state
func ResetOsChtimesWrapper() {
	Chtimes = os.Chtimes
}
