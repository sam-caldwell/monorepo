package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestIsScanEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			ScanEnabled: true,
		},
	} // Create an instance of the Manifest struct with ScanEnabled set to true

	// Test checking the scan enabled state
	scanEnabled := manifest.IsScanEnabled()

	// Verify that the returned scan enabled state matches the expected value
	assert.True(t, scanEnabled, "Returned scan enabled state should be true")
}
