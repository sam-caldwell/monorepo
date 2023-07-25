//go:build linux
// +build linux

package systemrecon

/*
 * GetFamily ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * GetFamily() for Linux
 *
 * 	Return the operating system family string using /etc/os-release
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
)

// GetFamily2 - return the operating system family
func GetFamily() (output string, err error) {

	var info keyvalue.KeyValue

	if info, err = GetLinuxOsReleaseMap(); err != nil {
		return output, fmt.Errorf(errors.UnsupportedOpsys)
	}

	if family, err := info.GetString("ID"); err == nil {
		return family, nil
	}

	return output, fmt.Errorf(errors.UnableToDetectFamily)
}
