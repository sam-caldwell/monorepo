package os

import (
	"os"
)

// Getppid - Abstract os.Getppid
var Getppid = ResetOsGetppidWrapper()

// ResetOsGetppidWrapper - Reset our os.Getppid wrapper to its original native state
func ResetOsGetppidWrapper() func() int {
	Getppid = os.Getppid
	return Getppid
}
