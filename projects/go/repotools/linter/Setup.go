package repolinter

/*
 * projects/repotools/repoLinter/Setup.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file installs linter dependencies and
 * any system configuration needed by them.
 */

import (
	packageManager2 "github.com/sam-caldwell/go/v2/projects/go/packageManager"
	"github.com/sam-caldwell/go/v2/projects/go/simpleLogger"
)

// Setup - install the repoLinter dependencies.
func Setup(logf simpleLogger.Logf, noop bool) (err error) {
	//on a noop, just return.  do nothing.  noop...no operation
	if !noop {
		packageList := []packageManager2.DependencyDescriptor{
			{
				Name:           "shellcheck",
				Installer:      packageManager2.Pkg,
				Detail:         "shellcheck",
				SkipIfDetected: true,
			},
			{
				Name:           "yamllint",
				Installer:      packageManager2.Pkg,
				Detail:         "yamllint",
				SkipIfDetected: true,
			},
			{
				Name:           "jsonlint",
				Installer:      packageManager2.Pkg,
				Detail:         "jsonlint",
				SkipIfDetected: true,
			},
			{
				Name:           "staticcheck_tools",
				Installer:      packageManager2.GoGet,
				Detail:         "honnef.co/go/tools/cmd/staticcheck",
				SkipIfDetected: true,
			},
			{
				Name:           "staticcheck",
				Installer:      packageManager2.GoInstall,
				Detail:         "honnef.co/go/tools/cmd/staticcheck@latest",
				SkipIfDetected: false,
			},
		}
		err = packageManager2.InstallDependencies(logf, packageList)
	}
	return err
}
