package projectmanifest

import (
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestSupportsOpsys(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Add supported platforms to the manifest
	manifest.SupportedPlatforms = []ManifestPlatforms{
		{Opsys: "darwin", CpuArch: "amd64"},
		{Opsys: "linux", CpuArch: "amd64"},
	}

	// Verify that the manifest supports the given operating system
	assert2.True(t, manifest.SupportsOpsys("darwin"), "Manifest should support 'darwin' operating system")
	assert2.False(t, manifest.SupportsOpsys("windows"), "Manifest should not support 'windows' operating system")
}
