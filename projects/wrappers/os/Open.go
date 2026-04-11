package os

import (
	"os"
)

// Open - Abstract os.Open
var Open = os.Open

// ResetOsOpenWrapper - Reset our os.Open wrapper to its original native state
func ResetOsOpenWrapper() {
	Open = os.Open
}
