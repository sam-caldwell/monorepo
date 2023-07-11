package os

import (
	"os"
)

// NewSyscallError - Abstract os.NewSyscallError
var NewSyscallError = os.NewSyscallError

// ResetOsNewSyscallErrorWrapper - Reset our os.NewSyscallError wrapper to its original native state
func ResetOsNewSyscallErrorWrapper() {
	NewSyscallError = os.NewSyscallError
}
