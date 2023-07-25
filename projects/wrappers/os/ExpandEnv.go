package os

import (
	"os"
)

// ExpandEnv - Abstract os.ExpandEnv
var ExpandEnv = os.ExpandEnv

// ResetOsExpandEnvWrapper - Reset our os.ExpandEnv wrapper to its original native state
func ResetOsExpandEnvWrapper() {
	ExpandEnv = os.ExpandEnv
}
