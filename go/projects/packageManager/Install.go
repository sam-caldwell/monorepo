package packageManager

/*
 * projects/packageManager/install.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file is the routing logic that calls a
 * specific class of installer based on the
 * DependencyDescriptor then returns the output
 * of executing the same.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
	"github.com/sam-caldwell/monorepo/go/projects/runcommand"
)

// Install - Detect the system's package manager and install the given package
func Install(thisPackage DependencyDescriptor) (out string, err error) {

	switch thisPackage.Installer {
	case GoGet:
		out, err = goGet(thisPackage)

	case GoInstall:
		out, err = goInstall(thisPackage)

	case Pkg:
		var pkgManager pkgManagerFunc
		if pkgManager, err = Detect(); err == nil {
			out, err = pkgManager(thisPackage)
		}

	case Shell:
		out, err = runcommand.ShellExecute(thisPackage.Detail)

	default:
		out = words.EmptyString
		err = fmt.Errorf("unsupported installer: %s", thisPackage.Installer.String())
	}
	return out, err
}
