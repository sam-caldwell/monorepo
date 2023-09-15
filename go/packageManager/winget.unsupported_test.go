//go:build !windows
// +build !windows

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"testing"
)

func TestWinget(t *testing.T) {
	out, err := winget(DependencyDescriptor{})
	if out != words.EmptyString {
		t.Fatal("output mismatch")
	}
	if err != nil {
		t.Fatal("error mismatch")
	}
}