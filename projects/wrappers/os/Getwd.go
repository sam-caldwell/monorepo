package os

import (
	"os"
)

// Getwd - Abstract os.Getwd
var Getwd = ResetOsGetwdWrapper()

// ResetOsGetwdWrapper - Reset our os.Getwd wrapper to its original native state
func ResetOsGetwdWrapper() func() (dir string, err error) {
	Getwd = os.Getwd
	return Getwd
}
