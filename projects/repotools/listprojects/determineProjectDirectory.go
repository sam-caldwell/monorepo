package listprojects

import "path/filepath"

// determineProjectDirectory - if determine which types of projects we will list.
func determineProjectDirectory(rootDirectory string, showCommands bool) (projectType string, numberOfParts int) {
	if showCommands {
		projectType, numberOfParts = "cmd", 4
	} else {
		projectType, numberOfParts = "projects", 3
	}
	return filepath.Join(rootDirectory, projectType), numberOfParts
}
