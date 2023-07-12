package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestSetName(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	name := "MyProject"

	// Set the project name
	manifest.SetName(name)

	// Verify that the project name is set correctly in the internal state
	assert.Equal(t, name, manifest.Name, "Project name should match the expected value")
}
