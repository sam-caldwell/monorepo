//go:build linux
// +build linux

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
)

/*
 * GetVersion ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * GetVersion() for Windows
 *
 * 	Return the operating system version string
 */

// GetVersion - return operating system version
func GetVersion() (output string, err error) {

	var info keyvalue.KeyValue

	if info, err = GetLinuxOsReleaseMap(); err != nil {
		return output, fmt.Errorf(errors.UnsupportedOpsys)
	}

	if version, err := info.GetString("VERSION_ID"); err == nil {
		return version, nil
	}

	return output, fmt.Errorf(errors.UnsupportedOpsys)
}
