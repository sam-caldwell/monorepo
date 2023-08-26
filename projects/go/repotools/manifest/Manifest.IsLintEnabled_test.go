package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestIsLintEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			LintEnabled: true,
		},
	} // Create an instance of the Manifest struct with LintEnabled set to true

	// Test checking the lint enabled state
	lintEnabled := manifest.IsLintEnabled()

	// Verify that the returned lint enabled state matches the expected value
	assert.True(t, lintEnabled, "Returned lint enabled state should be true")
}
