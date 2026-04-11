package os

import (
	"os"
)

// IsNotExist - Abstract os.IsNotExist
var IsNotExist = os.IsNotExist

// ResetOsIsNotExistWrapper - Reset our os.IsNotExist wrapper to its original native state
func ResetOsIsNotExistWrapper() {
	IsNotExist = os.IsNotExist
}
