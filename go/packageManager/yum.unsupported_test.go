//go:build !linux
// +build !linux

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"testing"
)

func TestYum(t *testing.T) {
	out, err := yum(DependencyDescriptor{})
	if out != words.EmptyString {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
