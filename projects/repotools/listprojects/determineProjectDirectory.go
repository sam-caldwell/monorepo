package listprojects

/*
 * projects/repotools/listprojects/determineProjectDirectory.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides a function used to determine the project directory
 * based on the repository root directory (rootDirectory) and a boolean
 * showCommands flag which determines if we should show our "command
 * projects" or our reusable code projects.
 */

import (
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"path/filepath"
)

// determineProjectDirectory - Determine type of projects to list and return its directory and dimensions
func determineProjectDirectory(rootDirectory string, showCommands bool) (projectDirectory string, numberOfParts int) {
	if showCommands {
		projectDirectory, numberOfParts = monorepo.CommandDirectory, monorepo.NumberOfPartsInCommandPath
	} else {
		projectDirectory, numberOfParts = monorepo.ProjectsDirectory, monorepo.NumberOfPartsInProjectPath
	}
	return filepath.Join(rootDirectory, projectDirectory), numberOfParts
}
