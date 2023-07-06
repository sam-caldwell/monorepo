package repotools

import (
	"github.com/sam-caldwell/go/v2/projects/packageManager"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
)

// Setup - Setup the common dependencies and configuration for repotools
func Setup(logf simpleLogger.Logf, noop bool) (err error) {
	//on a noop, just return.  do nothing.  noop...no operation
	if !noop {
		var packageList []packageManager.DependencyDescriptor
		err = packageManager.InstallDependencies(logf, packageList)
	}
	return err
}
