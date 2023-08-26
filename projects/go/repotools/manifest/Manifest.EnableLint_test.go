package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestEnableLint(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set LintEnabled to false
	manifest.Options.LintEnabled = false

	// Test enabling lint
	manifest.EnableLint()

	// Verify that LintEnabled is set to true
	assert.True(t, manifest.Options.LintEnabled, "LintEnabled should be set to true after calling EnableLint")
}
