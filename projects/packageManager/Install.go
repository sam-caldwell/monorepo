package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

// Install - Detect the system's package manager and install the given package
func Install(thisPackage DependencyDescriptor) (out string, err error) {
	var pkgManager pkgManagerFunc

	switch thisPackage.Installer {
	case GoGet:
		out, err = goGet(thisPackage)

	case GoInstall:
		out, err = goInstall(thisPackage)

	case Pkg:
		if pkgManager, err = Detect(); err == nil {
			out, err = pkgManager(thisPackage)
		}

	case Shell:
		out, err = runcommand.ShellExecute(thisPackage.Detail)

	default:
		out = words.EmptyString
		err = fmt.Errorf("unsupported installer: %s",
			thisPackage.Installer.String())
	}
	return out, err
}
