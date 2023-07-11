package repoTester

/*
 * projects/repotools/repoTester/Setup.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file installs repoTester dependencies and
 * any system configuration needed by them.
 */

import (
	"github.com/sam-caldwell/go/v2/projects/packageManager"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
)

// Setup - Install dependencies and configure repoTester
func Setup(logf simpleLogger.Logf, noop bool) (err error) {
	//on a noop, just return.  do nothing.  noop...no operation
	if !noop {
		var packageList []packageManager.DependencyDescriptor
		err = packageManager.InstallDependencies(logf, packageList)
	}
	return err
}
