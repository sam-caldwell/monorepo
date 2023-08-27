package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestIsSignEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			SignEnabled: true,
		},
	} // Create an instance of the Manifest struct with SignEnabled set to true

	// Test checking the sign enabled state
	signEnabled := manifest.IsSignEnabled()

	// Verify that the returned sign enabled state matches the expected value
	assert.True(t, signEnabled, "Returned sign enabled state should be true")
}
