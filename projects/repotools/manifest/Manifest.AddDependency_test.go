package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestAddDependency(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Test adding a dependency to the manifest
	dependency := "dependency-project"
	manifest = manifest.AddDependency(dependency)

	// Verify that the dependency is added to the Dependencies list
	assert.Equal(t, []string{dependency}, manifest.Dependencies, "Dependency should be added to the Dependencies list")
}
