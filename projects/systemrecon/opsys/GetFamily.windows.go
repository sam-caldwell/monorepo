//go:build windows
// +build windows

package systemrecon

import (
	"fmt"
)

func GetFamily(major, minor int) (version string, err error) {
	if major == 10 {
		version = "Windows 10"
		if minor != 0 {
			version = fmt.Sprintf("%s.%s", version, minor), nil
		}
		return version, err
	} else if major == 6 {
		switch minor {
		case 0:
			version = "Windows Vista"
		case 1:
			version = "Windows 7"
		case 2:
			version = "Windows 8"
		case 3:
			version = "Windows 8.1"
		default:
			version = fmt.Sprintf("%d.%d", major, minor)
			err = fmt.Errorf("unsupported minor version")
		}
		return version, err
	} else if major == 5 {
		switch minor {
		case 0:
			return "Windows 2000", nil
		case 1:
			return "Windows XP", nil
		case 2:
			return "Windows Server 2003", nil
		default:
			version = fmt.Sprintf("%d.%d", major, minor)
			err = fmt.Errorf("unsupported minor version")
		}
	} else {
		version = fmt.Sprintf("%d.%d", major, minor)
		err = fmt.Errorf("unsupported major version")
	}
	return version, err
}
