package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestIsTestEnabled(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			TestEnabled: true,
		},
	}
	testEnabled := manifest.IsTestEnabled()
	assert.True(t, testEnabled, "Returned TestEnabled state should be true")
}
