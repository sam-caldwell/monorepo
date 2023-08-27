package filters

import (
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestFilter_DefaultValues(t *testing.T) {
	filter := Filter{}
	assert2.False(t, filter.BuildEnabled, "buildEnabled should be false")
	assert2.False(t, filter.LintEnabled, "LintEnabled should be false")
	assert2.False(t, filter.PackEnabled, "PackEnabled should be false")
	assert2.False(t, filter.ScanEnabled, "ScanEnabled should be false")
	assert2.False(t, filter.SignEnabled, "SignEnabled should be false")
	assert2.Equal(t, words.EmptyString, filter.Language, "filter.Language should be empty")
	assert2.Equal(t, words.EmptyString, filter.os, "filter.os should be empty")
	assert2.Equal(t, words.EmptyString, filter.arch, "filter.arch should be empty")
}
