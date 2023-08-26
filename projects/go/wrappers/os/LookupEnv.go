package os

import (
	"os"
)

// LookupEnv - Abstract os.Hostname
var LookupEnv = os.LookupEnv

// ResetOsLookupEnvWrapper - Reset our os.LookupEnv wrapper to its original native state
func ResetOsLookupEnvWrapper() {
	LookupEnv = os.LookupEnv
}
