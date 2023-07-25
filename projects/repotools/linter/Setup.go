package repolinter

/*
 * projects/repotools/repoLinter/Setup.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file installs linter dependencies and
 * any system configuration needed by them.
 */

import (
	"github.com/sam-caldwell/go/v2/projects/packageManager"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
)

// Setup - install the repoLinter dependencies.
func Setup(logf simpleLogger.Logf, noop bool) (err error) {
	//on a noop, just return.  do nothing.  noop...no operation
	if !noop {
		packageList := []packageManager.DependencyDescriptor{
			{
				Name:           "shellcheck",
				Installer:      packageManager.Pkg,
				Detail:         "shellcheck",
				SkipIfDetected: true,
			},
			{
				Name:           "yamllint",
				Installer:      packageManager.Pkg,
				Detail:         "yamllint",
				SkipIfDetected: true,
			},
			{
				Name:           "jsonlint",
				Installer:      packageManager.Pkg,
				Detail:         "jsonlint",
				SkipIfDetected: true,
			},
			{
				Name:           "staticcheck_tools",
				Installer:      packageManager.GoGet,
				Detail:         "honnef.co/go/tools/cmd/staticcheck",
				SkipIfDetected: true,
			},
			{
				Name:           "staticcheck",
				Installer:      packageManager.GoInstall,
				Detail:         "honnef.co/go/tools/cmd/staticcheck@latest",
				SkipIfDetected: false,
			},
		}
		err = packageManager.InstallDependencies(logf, packageList)
	}
	return err
}
