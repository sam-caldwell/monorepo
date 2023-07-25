package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestGetName(t *testing.T) {
	manifest := &Manifest{
		Name: "MyProject",
	} // Create an instance of the Manifest struct with a specific project name

	// Test getting the project name
	name := manifest.GetName()

	// Verify that the returned project name matches the expected project name
	assert.Equal(t, "MyProject", name, "Returned project name should match the expected project name")
}
