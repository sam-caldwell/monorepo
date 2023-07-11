package os

import (
	"os"
)

// Expand - Abstract os.Expand
var Expand = ResetOsExpandWrapper()

// ResetOsExpandWrapper - Reset our os.Expand wrapper to its original native state
func ResetOsExpandWrapper() func(s string, mapping func(string) string) string {
	Expand = os.Expand
	return Expand
}
