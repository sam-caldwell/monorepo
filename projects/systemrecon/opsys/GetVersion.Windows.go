//go:build windows
// +build windows

package systemrecon

/*
 * GetVersion ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * GetVersion() for Windows
 *
 * 	Return the operating system version string
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"golang.org/x/sys/windows/registry"
	"runtime"
)

// GetVersion - return operating system version
func GetVersion() (version string, err error) {
	if runtime.GOOS != words.Windows {
		return words.EmptyString, fmt.Errorf(errors.UnsupportedOpsys)
	}
	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE)

	if err != nil {
		return words.EmptyString, err
	}

	defer func() { _ = k.Close() }()

	majVersion, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return words.EmptyString, err
	}

	minVersion, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		return words.EmptyString, err
	}

	return fmt.Sprintf("%d.%d", majVersion, minVersion), nil
}
