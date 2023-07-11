package os

import (
	"os"
)

// IsExist - Abstract os.IsExist
var IsExist = ResetOsIsExistWrapper()

// ResetOsIsExistWrapper - Reset our os.IsExist wrapper to its original native state
func ResetOsIsExistWrapper() func(err error) bool {
	IsExist = os.IsExist
	return IsExist
}
