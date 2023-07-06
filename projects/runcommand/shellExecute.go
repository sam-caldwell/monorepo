//go:build !windows
// +build !windows

package runcommand

import (
	"os/exec"
)

// ShellExecute - A simple shell executor for windows or non-windows systems
func ShellExecute(command string) (out string, err error) {
	var raw []byte
	raw, err = exec.Command("sh", "-c", command).Output()
	return string(raw), err
}
