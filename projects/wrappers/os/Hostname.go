package os

import (
	"os"
)

// Hostname - Abstract os.Hostname
var Hostname = ResetOsHostnameWrapper()

// ResetOsHostnameWrapper - Reset our os.Hostname wrapper to its original native state
func ResetOsHostnameWrapper() func() (name string, err error) {
	Hostname = os.Hostname
	return Hostname
}
