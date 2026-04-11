package os

import (
	"os"
)

// IsExist - Abstract os.IsExist
var IsExist = os.IsExist

// ResetOsIsExistWrapper - Reset our os.IsExist wrapper to its original native state
func ResetOsIsExistWrapper() {
	IsExist = os.IsExist

}
