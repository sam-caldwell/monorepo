//go:build !windows
// +build !windows

package packageManager

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/misc/words"
	"testing"
)

func TestChoco(t *testing.T) {
	out, err := choco(DependencyDescriptor{})
	if out != words.EmptyString {
		t.Fatal("output mismatch")
	}
	if err != nil {
		t.Fatal("error mismatch")
	}
}
