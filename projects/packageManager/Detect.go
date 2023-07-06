package packageManager

/*
 * projects/packageManager/Detect.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines a function which will detect any installed and
 * supported package manager using a map of package managers and
 * their supported operating systems.
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"runtime"
)

// Detect - Detect if the system has a supported package manager
func Detect() (pkg pkgManagerFunc, err error) {
	/*
	 * Create map of our options
	 */
	var packageManagerList = map[string]map[string]pkgManagerFunc{
		"windows": {
			"winget": winget,
			"choco":  choco,
		},
		"ubuntu": {
			"apt-get": aptGet,
			"brew":    brew,
		},
		"centos": {
			"yum":  aptGet,
			"brew": brew,
		},
		"fedora": {
			"yum":  aptGet,
			"brew": brew,
		},
		"redhat": {
			"yum":  aptGet,
			"brew": brew,
		},
		"darwin": {
			"brew": brew,
		},
	}
	switch opsys := runtime.GOOS; opsys {
	case words.Darwin:
		return findSupportedPackageManager(packageManagerList[words.Darwin])

	case words.Linux:
		switch distribution := detectLinuxDistribution(); distribution {
		case words.Ubuntu, words.Debian:
			return findSupportedPackageManager(packageManagerList[words.Ubuntu])

		case words.Redhat, words.Centos, words.Fedora:
			return findSupportedPackageManager(packageManagerList[words.Redhat])

		default:
			return nil, fmt.Errorf(errors.UnsupportedOpsys+errors.Details, distribution)
		} /*end of switch*/

	case words.Windows:
		return findSupportedPackageManager(packageManagerList[words.Windows])

	default:
		return nil, fmt.Errorf(errors.UnsupportedOpsys+errors.Details, opsys)
	} /*end of switch*/
}
