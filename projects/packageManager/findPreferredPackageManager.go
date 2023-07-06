package packageManager

import (
	"fmt"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
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
