package packageManager

/*
 * projects/packageManager/findSupportedPackageManager.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file creates a function which checks the system to see
 * if a given package manager is installed and accessible given
 * a map of package managers supported on the system.  It will
 * return a function pointer referencing the package manager's
 * wrapper function.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/systemrecon"
)

// findSupportedPackageManager - Find out if the system has a supported package manager
func findSupportedPackageManager(pkgMgrList map[string]pkgManagerFunc) (pkgManagerFunc, error) {
	/*
	 * loop through the package manager options and terminate
	 * on the first finding.  return an error if no suitable
	 * match exists.
	 */
	for name, pkgMgrFunction := range pkgMgrList {
		if exitcode, _ := systemrecon.HasExecutable(name); exitcode == 0 {
			return pkgMgrFunction, nil
		}
	}
	return nil, fmt.Errorf("no package manager detected")
}
