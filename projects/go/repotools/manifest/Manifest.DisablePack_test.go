package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestDisablePack(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set PackEnabled to true
	manifest.Options.PackEnabled = true

	// Test disabling pack
	manifest.DisablePack()

	// Verify that PackEnabled is set to false
	assert.False(t, manifest.Options.PackEnabled, "PackEnabled should be set to false after calling DisablePack")
}
