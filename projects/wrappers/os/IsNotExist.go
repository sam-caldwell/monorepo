package os

import (
	"os"
)

// IsNotExist - Abstract os.IsNotExist
var IsNotExist = ResetOsIsNotExistWrapper()

// ResetOsIsNotExistWrapper - Reset our os.IsNotExist wrapper to its original native state
func ResetOsIsNotExistWrapper() func(err error) bool {
	IsNotExist = os.IsNotExist
	return IsNotExist
}
