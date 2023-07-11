package os

import (
	"os"
)

// Exit - Abstract os.Exit
var Exit = ResetOsExitWrapper()

// ResetOsExitWrapper - Reset our os.Exit wrapper to its original native state
func ResetOsExitWrapper() func(n int) {
	Exit = os.Exit
	return Exit
}
