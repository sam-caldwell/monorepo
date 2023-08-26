package packageManager

/*
 * projects/packageManager/InstallDependencies.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file is the package manager package top-level
 * function for installing arbitrary dependencies given
 * a list of DependencyDescriptor records. Each record
 * defines the dependency, how it is to be installed
 * and other characteristics.
 *
 * This function will install the dependency and then
 * log the result of the operation to the UI for the
 * user via simpleLogger.
 */

import (
	ansi "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/simpleLogger"
)

// InstallDependencies - Install dependencies for the caller
func InstallDependencies(logf simpleLogger.Logf, packageList []DependencyDescriptor) (err error) {

	count := len(packageList)

	for n, thisPackage := range packageList {

		logf(ansi.Blue(), "installing (%d,%d) %s", n, count, thisPackage.Name)

		if out, err := Install(thisPackage); err != nil {

			logf(ansi.Red(), "Installation failed on %s\n"+
				"Details:\n"+
				"%s\n", thisPackage.Name, out)

			return err
		}

		logf(ansi.Green(), "installation OK (%d,%d) %s", n, count, thisPackage.Name)

	}
	return nil
}
