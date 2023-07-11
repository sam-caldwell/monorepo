package os

import (
	"os"
)

// Open - Abstract os.Open
var Open = ResetOsOpenWrapper()

// ResetOsOpenWrapper - Reset our os.Open wrapper to its original native state
func ResetOsOpenWrapper() func(name string) (*os.File, error) {
	Open = os.Open
	return Open
}
