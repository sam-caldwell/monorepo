package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestSupportsArch(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Add supported platforms to the manifest
	manifest.SupportedPlatforms = []ManifestPlatforms{
		{Opsys: "darwin", CpuArch: "amd64"},
		{Opsys: "linux", CpuArch: "amd64"},
	}

	// Verify that the manifest supports the given CPU architecture
	assert.True(t, manifest.SupportsArch("amd64"), "Manifest should support 'amd64' architecture")
	assert.False(t, manifest.SupportsArch("arm64"), "Manifest should not support 'arm64' architecture")
}
