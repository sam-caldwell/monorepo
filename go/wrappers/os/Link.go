package os

import (
	"os"
)

// Link - Abstract os.Link
var Link = os.Link

// ResetOsLinkWrapper - Reset our os.Link wrapper to its original native state
func ResetOsLinkWrapper() {
	Link = os.Link
}
