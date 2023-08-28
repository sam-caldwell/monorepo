//go:build darwin
// +build darwin

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/monorepo/v2/projects/go/misc/words"
	"os/exec"
	"strconv"
	"strings"
)

/*
 * SystemInfo ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * SystemInfo() for Darwin
 *
 * 	Return the amount of RAM in the system (in KB)
 */

// RamSize - Return the ram size in KB
func RamSize() (int, error) {
	cmd := exec.Command(words.MacSysCtl, words.MacSysCtlHwMemSize)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	outputStr := strings.TrimSpace(string(output))
	fields := strings.Fields(outputStr)
	if len(fields) < 2 {
		return 0, fmt.Errorf(errors.InternalError)
	}

	size, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, err
	}

	// Convert from bytes to kilobytes and return
	return size / 1024, nil
}
