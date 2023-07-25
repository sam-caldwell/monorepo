package os

import (
	"os"
)

// Getwd - Abstract os.Getwd
var Getwd = os.Getwd

// ResetOsGetwdWrapper - Reset our os.Getwd wrapper to its original native state
func ResetOsGetwdWrapper() {
	Getwd = os.Getwd
}
