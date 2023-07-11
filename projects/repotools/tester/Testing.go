package repoTester

import "github.com/sam-caldwell/go/v2/projects/simpleLogger"

// Testing - Run all tests for the repo or specific project
func Testing(logf simpleLogger.Logf, noop bool, projectPath string) error {
	if noop {
		return nil
	}
	return nil
}
