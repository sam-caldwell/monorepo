package projectmanifest

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"os"
	"strings"
	"testing"
)

func TestWriteFile(t *testing.T) {
	manifest := &Manifest{
		fileName: "test.yaml",
		Name:     "test_project",
		Author:   "Sam Caldwell",
		Options: ManifestOptions{
			BuildEnabled: false,
			LintEnabled:  true,
			PackEnabled:  false,
			ScanEnabled:  true,
			SignEnabled:  false,
			TestEnabled:  true,
			Language:     "go",
		},
		SupportedPlatforms: []ManifestPlatforms{
			{Opsys: "darwin", CpuArch: "amd64"},
			{Opsys: "linux", CpuArch: "amd64"},
		},
		Dependencies: []string{"ansi", "testing"},
	}

	// Call the WriteFile method to write the manifest to a file
	manifest.WriteFile()

	// Read the content of the written file
	content, err := os.ReadFile(manifest.fileName)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	// Verify that the content matches the expected YAML representation
	expectedContent := []string{
		"name: test_project",
		"author: Sam Caldwell",
		"options:",
		"    buildEnabled: false",
		"    lintEnabled: true",
		"    packEnabled: false",
		"    scanEnabled: true",
		"    signEnabled: false",
		"    testEnabled: true",
		"    language: go",
		"supportedPlatforms:",
		"    - opsys: darwin",
		"      arch: amd64",
		"    - opsys: linux",
		"      arch: amd64",
		"dependencies:",
		"    - ansi",
		"    - testing",
		""}
	for lineNo, line := range strings.Split(string(content), "\n") {
		assert.Equal(t, line, expectedContent[lineNo], fmt.Sprintf("File content failed to match at line %d: '%s' <> '%s'", lineNo, line, expectedContent[lineNo]))
	}

	// Clean up the test file
	err = os.Remove(manifest.fileName)
	if err != nil {
		t.Fatalf("failed to remove test file: %v", err)
	}
}
