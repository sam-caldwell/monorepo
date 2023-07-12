package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestAddPlatform(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Test adding a supported platform to the manifest
	operatingSystem := "linux"
	cpuArchitecture := "amd64"
	manifest = manifest.AddPlatform(operatingSystem, cpuArchitecture)

	// Verify that the platform is added to the SupportedPlatforms list
	expectedPlatforms := []ManifestPlatforms{
		{
			Opsys:   operatingSystem,
			CpuArch: cpuArchitecture,
		},
	}
	assert.Equal(t, expectedPlatforms, manifest.SupportedPlatforms, "Platform should be added to the SupportedPlatforms list")

	// Test adding an unsupported operating system
	unsupportedOS := "solaris"
	manifest = manifest.AddPlatform(unsupportedOS, cpuArchitecture)
	assert.NotNil(t, manifest.err, "Error should be set for unsupported operating system")
	assert.Equal(t, errors.UnsupportedOpsys, manifest.err.Error(), "Error message should indicate unsupported operating system")

	// Test adding an unsupported CPU architecture
	unsupportedArch := "arm"
	manifest = manifest.AddPlatform(operatingSystem, unsupportedArch)
	assert.NotNil(t, manifest.err, "Error should be set for unsupported CPU architecture")
	assert.Equal(t, errors.UnsupportedCpuArchitecture, manifest.err.Error(), "Error message should indicate unsupported CPU architecture")
}
