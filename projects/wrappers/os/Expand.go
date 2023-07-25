package os

import (
	"os"
)

// Expand - Abstract os.Expand
var Expand = os.Expand

// ResetOsExpandWrapper - Reset our os.Expand wrapper to its original native state
func ResetOsExpandWrapper() {
	Expand = os.Expand
}
