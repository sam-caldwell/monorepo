package packageManager

import (
	"testing"
)

func TestGoInstall(t *testing.T) {
	// Test case 1: Package installation succeeds
	pkg := DependencyDescriptor{
		Name:   "honnef.co/go/tools/cmd/staticcheck@latest",
		Detail: "honnef.co/go/tools/cmd/staticcheck@latest",
	}

	_, err := goInstall(pkg)

	if err != nil {
		t.Fatal("error mismatch")
	}
}
