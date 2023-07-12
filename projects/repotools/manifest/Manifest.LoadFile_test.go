package projectmanifest

import (
	"os"
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestLoadFile(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	fileName := "/tmp/test_manifest.yaml"

	// Create a temporary YAML file for testing
	const yamlData = `---
name: "test_project"
author: "Sam Caldwell"
options:
  buildEnabled: false
  lintEnabled: true
  packEnabled: false
  scanEnabled: true
  signEnabled: false
  language: go
supportedPlatforms:
  - opsys: darwin
    arch: amd64
  - opsys: darwin
    arch: arm64
  - opsys: linux
    arch: amd64
  - opsys: linux
    arch: arm64
  - opsys: windows
    arch: amd64
dependencies:
  - ansi
  - testing
`

	err := os.WriteFile(fileName, []byte(yamlData), 0644)
	assert.Nil(t, err, "Error writing test manifest file")
	defer func() { _ = os.Remove(fileName) }() // Clean up the temporary file

	// Load the manifest file
	err = manifest.LoadFile(fileName)

	// Verify that no error occurred
	assert.Nil(t, err, "Error loading manifest file")

	// Verify the loaded values
	assert.Equal(t, "go", manifest.Options.Language, "Loaded language should match expected value")
	assert.True(t, manifest.Options.LintEnabled, "Loaded lint enabled state should be true")
	assert.True(t, manifest.Options.ScanEnabled, "Loaded scan enabled state should be true")
	assert.False(t, manifest.Options.BuildEnabled, "Loaded build enabled state should be false")
	assert.False(t, manifest.Options.PackEnabled, "Loaded pack enabled state should be false")
	assert.False(t, manifest.Options.SignEnabled, "Loaded sign enabled state should be false")
	assert.Equal(t, 5, len(manifest.SupportedPlatforms), "Loaded supported platforms should contain 5 entries")
	assert.Equal(t, 2, len(manifest.Dependencies), "Loaded dependencies should contain 2 entries")
}
