package packageManager

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
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
