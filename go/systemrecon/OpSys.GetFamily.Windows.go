//go:build windows
// +build windows

package systemrecon

/*
 * GetFamily ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * GetFamily() for Windows
 *
 * 	Return the operating system family string
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/semver"
)

func GetFamily() (result string, err error) {
	var version semver.SemanticVersion
	{
		rawVersion, err := GetVersion()
		if err != nil {
			return rawVersion, err
		}
		if err = version.ParseP(&rawVersion); err != nil {
			fmt.Println("raw:", rawVersion, ":", err.Error())
			return version.String(), err
		}
	}
	switch version.Major() {
	case 10:
		result = words.Windows10
		if version.Minor() != 0 {
			result = fmt.Sprintf("%d.%d", version.Major(), version.Minor())
		}
		return result, err
	case 6:
		switch version.Minor() {
		case 0:
			result = words.WindowsVista
		case 1:
			result = words.Windows7
		case 2:
			result = words.Windows8
		case 3:
			result = words.Windows81
		default:
			result = fmt.Sprintf("%d.%d", version.Major(), version.Minor())
			err = fmt.Errorf(errors.UnsupportedVersion)
		}
		return result, err
	case 5:
		switch version.Minor() {
		case 0:
			result = words.Windows2K
		case 1:
			result = words.WindowsXp
		case 2:
			result = words.WindowsServer2K3
		default:
			result = fmt.Sprintf("%d.%d", version.Major(), version.Minor())
			err = fmt.Errorf(errors.UnsupportedVersion)
		}
	default:
		result = fmt.Sprintf("%d.%d", version.Major(), version.Minor())
		err = fmt.Errorf(errors.UnsupportedVersion)
	}
	return result, err
}
