package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestEnableBuild(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set BuildEnabled to false
	manifest.Options.BuildEnabled = false

	// Test enabling build
	manifest.EnableBuild()

	// Verify that BuildEnabled is set to true
	assert.True(t, manifest.Options.BuildEnabled, "BuildEnabled should be set to true after calling EnableBuild")
}
