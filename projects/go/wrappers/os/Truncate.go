package os

import (
	"os"
)

// Truncate - Abstract os.Truncate
var Truncate = os.Truncate

// ResetOsTruncateWrapper - Reset our os.Truncate wrapper to its original native state
func ResetOsTruncateWrapper() {
	Truncate = os.Truncate
}
