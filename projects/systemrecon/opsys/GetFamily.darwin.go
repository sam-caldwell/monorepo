//go:build darwin
// +build darwin

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/semver"
)

// GetFamily - Return the Operating System family
func GetFamily() (family string, err error) {
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
		switch version.Minor() {
		case 7:
			family = words.OsxLion
		case 8:
			family = words.OsxMountainLion
		case 9:
			family = words.OsxMavericks
		case 10:
			family = words.OsxYosemite
		case 11:
			family = words.OsxElCapitan
		case 12:
			family = words.MacOsSierra
		case 13:
			family = words.MacOsHighSierra
		case 14:
			family = words.MacOsMojave
		case 15:
			family = words.MacOsCatalina
		}
	case 11:
		family = words.MacOsBigSur
	case 12:
		family = words.MacOsMonterey
	case 13:
		family = words.MacOsVentura
	case 14:
		family = words.MacOsSonoma
	default:
		family = version.String()
	}
	return family, err
}
