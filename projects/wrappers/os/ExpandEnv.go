package os

import (
	"os"
)

// ExpandEnv - Abstract os.ExpandEnv
var ExpandEnv = ResetOsExpandEnvWrapper()

// ResetOsExpandEnvWrapper - Reset our os.ExpandEnv wrapper to its original native state
func ResetOsExpandEnvWrapper() func(s string) string {
	ExpandEnv = os.ExpandEnv
	return ExpandEnv
}
