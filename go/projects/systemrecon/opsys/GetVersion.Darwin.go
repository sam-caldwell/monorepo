//go:build darwin
// +build darwin

package systemrecon

/*
 * GetVersion ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * GetVersion() for MacOs
 *
 * 	Return the operating system version string
 */

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
	"os/exec"
	"strings"
)

// GetVersion - return operating system version
func GetVersion() (version string, err error) {
	out, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return words.EmptyString, err
	}
	version = strings.TrimSpace(string(out))
	return version, nil
}
