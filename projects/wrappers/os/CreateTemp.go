package os

import (
	"os"
)

// CreateTemp - Abstract os.CreateTemp
var CreateTemp = ResetOsCreateTempWrapper()

// ResetOsCreateTempWrapper - Reset our os.CreateTemp wrapper to its original native state
func ResetOsCreateTempWrapper() func(dir, pattern string) (*os.File, error) {
	CreateTemp = os.CreateTemp
	return CreateTemp
}
