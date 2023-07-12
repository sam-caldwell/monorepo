package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestEnablePack(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set PackEnabled to false
	manifest.Options.PackEnabled = false

	// Test enabling pack
	manifest.EnablePack()

	// Verify that PackEnabled is set to true
	assert.True(t, manifest.Options.PackEnabled, "PackEnabled should be set to true after calling EnablePack")
}
