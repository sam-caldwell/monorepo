package os

import "os"

// Create - Abstract os.Create
var Create = os.Create

// ResetOsCreate - Reset our os.Create wrapper to its native state
func ResetOsCreate() {
	Create = os.Create
}
