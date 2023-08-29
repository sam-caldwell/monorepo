package os

import (
	"os"
)

// Exit - Abstract os.Exit
var Exit = os.Exit

// ResetOsExitWrapper - Reset our os.Exit wrapper to its original native state
func ResetOsExitWrapper() {
	Exit = os.Exit
}
