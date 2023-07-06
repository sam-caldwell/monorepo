package packageManager

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
)

// Install - Detect the system's package manager and install the given package
func Install(packageName string) (out string, err error) {
	var pkgManager pkgManagerFunc
	if pkgManager, err = Detect(); err != nil {
		return words.EmptyString, err
	}
	return pkgManager(packageName)
}
