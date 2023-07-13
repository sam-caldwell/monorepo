package listprojects_test

import (
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	"github.com/sam-caldwell/go/v2/projects/repotools/listprojects"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestListAsOrderedPair(t *testing.T) {
	expectedProjects := func() []string {
		var subdirectories []string

		rootDir, _ := repotools.FindRepoRoot()
		projectDir := filepath.Join(rootDir, "projects")
		files, err := os.ReadDir(projectDir)
		if err != nil {
			return []string{}
		}
		for _, thisFile := range files {
			if file.Exists(filepath.Join(projectDir, thisFile.Name(), "MANIFEST.yaml")) {
				if thisFile.IsDir() {
					subdirectories = append(subdirectories, thisFile.Name())
				}
			}
		}
		sort.Strings(subdirectories)
		return subdirectories
	}()
	var filter filters.Filter

	result, err := listprojects.ListAsOrderedPair(filter)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for i, actual := range result {
		if filepath.Base(actual.Key) != expectedProjects[i] {
			t.Fatalf("Mismatch at %d. expected %s git %s", i, expectedProjects[i], actual)
		}
	}
}
