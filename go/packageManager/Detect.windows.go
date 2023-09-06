//go:build windows
// +build windows

package packageManager

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
	var packageManagerList = map[string]pkgManagerFunc{
		"winget": winget,
		"choco":  choco,
	}
	return findSupportedPackageManager(packageManagerList)
}
