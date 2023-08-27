package projectmanifest

import (
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestSupportsArch(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Add supported platforms to the manifest
	manifest.SupportedPlatforms = []ManifestPlatforms{
		{Opsys: "darwin", CpuArch: "amd64"},
		{Opsys: "linux", CpuArch: "amd64"},
	}

	// Verify that the manifest supports the given CPU architecture
	assert2.True(t, manifest.SupportsArch("amd64"), "Manifest should support 'amd64' architecture")
	assert2.False(t, manifest.SupportsArch("arm64"), "Manifest should not support 'arm64' architecture")
}
