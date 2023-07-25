package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestDisableBuild(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set BuildEnabled to true
	manifest.Options.BuildEnabled = true

	// Test disabling build
	manifest.DisableBuild()

	// Verify that BuildEnabled is set to false
	assert.False(t, manifest.Options.BuildEnabled, "BuildEnabled should be set to false after calling DisableBuild")
}
