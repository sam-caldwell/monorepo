//go:build linux
// +build linux

package packageManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYum(t *testing.T) {
	// Test case 1: Package installation succeeds
	pkg := DependencyDescriptor{
		Name:           "package-name",
		Detail:         "package-name",
		SkipIfDetected: false,
	}

	output, err := yum(pkg)
	expectedOutput := "yum install -y package-name"

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)

	// Test case 2: Package installation skipped
	pkg = DependencyDescriptor{
		Name:           "existing-package",
		Detail:         "existing-package",
		SkipIfDetected: true,
	}

	output, err = yum(pkg)
	expectedOutput = "dependency exists. skipping"

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
