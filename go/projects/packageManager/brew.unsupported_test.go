//go:build !linux || !darwin
// +build !linux !darwin

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
	"testing"
)

func TestBrew(t *testing.T) {
	output, err := brew(DependencyDescriptor{})
	if output != words.EmptyString {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
