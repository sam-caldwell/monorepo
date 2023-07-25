package os

import (
	"os"
)

// Getppid - Abstract os.Getppid
var Getppid = os.Getppid

// ResetOsGetppidWrapper - Reset our os.Getppid wrapper to its original native state
func ResetOsGetppidWrapper() {
	Getppid = os.Getppid
}
