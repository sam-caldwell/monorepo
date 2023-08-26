package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestIsBuildEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			BuildEnabled: true,
		},
	} // Create an instance of the Manifest struct with BuildEnabled set to true

	// Test checking the build enabled state
	buildEnabled := manifest.IsBuildEnabled()

	// Verify that the returned build enabled state matches the expected value
	assert.True(t, buildEnabled, "Returned build enabled state should be true")
}
