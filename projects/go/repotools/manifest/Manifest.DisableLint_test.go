package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestDisableLint(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set LintEnabled to true
	manifest.Options.LintEnabled = true

	// Test disabling lint
	manifest.DisableLint()

	// Verify that LintEnabled is set to false
	assert.False(t, manifest.Options.LintEnabled, "LintEnabled should be set to false after calling DisableLint")
}
