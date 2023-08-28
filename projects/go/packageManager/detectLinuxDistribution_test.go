package packageManager

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Variable to hold the original implementation of os.ReadFile
var osReadFile = os.ReadFile

func TestDetectLinuxDistribution(t *testing.T) {
	// Test case 1: Linux distribution detected from mock /etc/os-release
	mockOSRelease := []byte(`ID=ubuntu
VERSION_ID=20.04
NAME="Ubuntu"`)

	expectedDistribution := "ubuntu"

	// Mocking the os.ReadFile function
	os.ReadFile = func(filename string) ([]byte, error) {
		return mockOSRelease, nil
	}
	defer func() {
		// Reset os.ReadFile to its original implementation after the test
		os.ReadFile = osReadFile
	}()

	opsys := detectLinuxDistribution()

	assert.Equal(t, expectedDistribution, opsys)
}
