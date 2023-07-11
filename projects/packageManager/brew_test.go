//go:build linux && darwin
// +build linux,darwin

package packageManager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrew(t *testing.T) {
	// Test case 1: Dependency does not exist
	pkg := DependencyDescriptor{
		Name:           "nonexistent-package",
		Detail:         "nonexistent-package",
		SkipIfDetected: true,
	}

	output, err := brew(pkg)
	assert.NoError(t, err)
	assert.Equal(t, "brew install nonexistent-package", output)

	// Test case 2: Dependency already exists, skip installation
	pkg = DependencyDescriptor{
		Name:           "existing-package",
		Detail:         "existing-package",
		SkipIfDetected: true,
	}

	output, err = brew(pkg)
	assert.NoError(t, err)
	assert.Equal(t, "dependency exists. skipping", output)

	// Test case 3: Dependency already exists, do not skip installation
	pkg = DependencyDescriptor{
		Name:           "existing-package",
		Detail:         "existing-package",
		SkipIfDetected: false,
	}

	output, err = brew(pkg)
	assert.NoError(t, err)
	assert.Equal(t, "brew install existing-package", output)
}
