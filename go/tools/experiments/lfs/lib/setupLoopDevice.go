package lfs

import (
	"fmt"
	"os/exec"
	"strings"
)

// setupLoopDevice - sets up a loop device for the given file.
func setupLoopDevice(fileName string) (string, error) {
	// Find an available loop device
	raw, err := exec.Command("losetup", "-f", "--show", fileName).Output()
	if err != nil {
		return "", fmt.Errorf("losetup command failed: %w", err)
	}

	// Trim the output to get the loop device path
	lines := strings.Split(string(raw), "\n")
	loopDevice := lines[0] // Remove the newline character

	return loopDevice, nil
}
