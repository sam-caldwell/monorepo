package listprojects

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/testing/assert"
	"path/filepath"
	"testing"
)

func TestDetermineProjectDirectory(t *testing.T) {
	const rootDirectory = "/fake/path/to/root"

	testFunc := func(testName string, showCommands bool, expectedProjectType string, expectedNumberOfParts int) {

		actualProjectType, actualNumberOfParts := determineProjectDirectory(rootDirectory, showCommands)

		assert.True(t, actualProjectType == expectedProjectType,
			fmt.Sprintf("%s: Expect projectType %s got %s", testName, expectedProjectType, actualProjectType))

		if actualNumberOfParts != expectedNumberOfParts {
			t.Fatalf("%s Expected numberOfParts to be %d, but got %d", testName, expectedNumberOfParts, actualNumberOfParts)
		}
	}

	testFunc("cmdHappy",
		true,
		filepath.Join(rootDirectory, "cmd"),
		4)

	testFunc("projectsHappy",
		false,
		filepath.Join(rootDirectory, "projects"),
		3)
}
