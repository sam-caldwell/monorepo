package os

import (
	"os"
)

// Hostname - Abstract os.Hostname
var Hostname = os.Hostname

// ResetOsHostnameWrapper - Reset our os.Hostname wrapper to its original native state
func ResetOsHostnameWrapper() {
	Hostname = os.Hostname
}
