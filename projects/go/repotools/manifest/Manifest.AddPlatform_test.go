package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
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
	assert2.Equal(t, expectedPlatforms, manifest.SupportedPlatforms, "Platform should be added to the SupportedPlatforms list")

	// Test adding an unsupported operating system
	unsupportedOS := "solaris"
	manifest = manifest.AddPlatform(unsupportedOS, cpuArchitecture)
	assert2.NotNil(t, manifest.err, "Error should be set for unsupported operating system")
	assert2.Equal(t, errors.UnsupportedOpsys, manifest.err.Error(), "Error message should indicate unsupported operating system")

	// Test adding an unsupported CPU architecture
	unsupportedArch := "arm"
	manifest = manifest.AddPlatform(operatingSystem, unsupportedArch)
	assert2.NotNil(t, manifest.err, "Error should be set for unsupported CPU architecture")
	assert2.Equal(t, errors.UnsupportedCpuArchitecture, manifest.err.Error(), "Error message should indicate unsupported CPU architecture")
}
