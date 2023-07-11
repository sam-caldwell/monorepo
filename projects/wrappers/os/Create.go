package os

import "os"

// Create - Abstract os.Create
var Create = ResetOsCreate()

// ResetOsCreate - Reset our os.Create wrapper to its native state
func ResetOsCreate() func(name string) (*os.File, error) {
	Create = os.Create
	return Create
}
