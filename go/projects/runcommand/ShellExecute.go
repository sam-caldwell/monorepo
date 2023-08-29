//go:build !windows
// +build !windows

package runcommand

/*
 * projects/runcommand/ShellExecute.go (non-windows)
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file executes an arbitrary shell command and
 * returns the stdout output and any error
 */
import (
	"os/exec"
)

// ShellExecute - A simple shell executor for windows or non-windows systems
func ShellExecute(command string) (out string, err error) {
	var raw []byte
	raw, err = exec.Command("sh", "-c", command).Output()
	return string(raw), err
}
