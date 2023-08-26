package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestManifestPlatforms(t *testing.T) {
	platform := ManifestPlatforms{} // Create an instance of the ManifestPlatforms struct

	// Test the structure of the ManifestPlatforms
	assert.Equal(t, "", platform.Opsys, "Default value of Opsys should be an empty string")
	assert.Equal(t, "", platform.CpuArch, "Default value of CpuArch should be an empty string")
}
