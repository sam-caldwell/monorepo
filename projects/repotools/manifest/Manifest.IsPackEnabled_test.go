package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestIsPackEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			PackEnabled: true,
		},
	} // Create an instance of the Manifest struct with PackEnabled set to true

	// Verify that the returned pack enabled state matches the expected value
	assert.True(t, manifest.IsPackEnabled(), "Returned pack enabled state should be true")
}
