//go:build windows
// +build windows

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"golang.org/x/sys/windows/registry"
	"runtime"
)

func GetVersion() (version string, err error) {
	if runtime.GOOS != words.Windows {
		return words.EmptyString, fmt.Error(exit.ErrUnsupportedOpsys)
	}
	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE)

	if err != nil {
		return "", err
	}

	defer func() { _ = k.Close() }()

	majVersion, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return "", err
	}

	minVersion, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%d", majVersion, minVersion), nil
}
