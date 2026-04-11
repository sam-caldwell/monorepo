package os

import "os"

// CreateTemp - Abstract os.CreateTemp
var CreateTemp = os.CreateTemp

// ResetOsCreateTempWrapper - Reset our os.CreateTemp wrapper to its original native state
func ResetOsCreateTempWrapper() {
	CreateTemp = os.CreateTemp
}
