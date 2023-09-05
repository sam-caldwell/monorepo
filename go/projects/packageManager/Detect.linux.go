//go:build linux
// +build linux

package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/exit/errors"
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
)

/*
 * projects/packageManager/Detect.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines a function which will detect any installed and
 * supported package manager using a map of package managers and
 * their supported operating systems.
 */

// Detect - Detect if the system has a supported package manager
func Detect() (pkg pkgManagerFunc, err error) {
	var packageManagerList = map[string]map[string]pkgManagerFunc{
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
	}
	switch distribution := detectLinuxDistribution(); distribution {
	case words.Ubuntu, words.Debian:
		pkg, err = findSupportedPackageManager(packageManagerList[words.Ubuntu])

	case words.Redhat, words.Centos, words.Fedora:
		pkg, err = findSupportedPackageManager(packageManagerList[words.Redhat])

	default:
		pkg, err = nil, fmt.Errorf(errors.UnsupportedOpsys+errors.Details, distribution)
	} /*end of switch*/
	return pkg, err
}
