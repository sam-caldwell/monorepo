package filters

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/testing/assert"
	"testing"
)

func TestFilter_DefaultValues(t *testing.T) {
	filter := Filter{}
	assert.False(t, filter.BuildEnabled, "buildEnabled should be false")
	assert.False(t, filter.LintEnabled, "LintEnabled should be false")
	assert.False(t, filter.PackEnabled, "PackEnabled should be false")
	assert.False(t, filter.ScanEnabled, "ScanEnabled should be false")
	assert.False(t, filter.SignEnabled, "SignEnabled should be false")
	assert.Equal(t, words.EmptyString, filter.Language, "filter.Language should be empty")
	assert.Equal(t, words.EmptyString, filter.os, "filter.os should be empty")
	assert.Equal(t, words.EmptyString, filter.arch, "filter.arch should be empty")
}
