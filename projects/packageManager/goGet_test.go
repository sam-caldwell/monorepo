package packageManager

import (
	"testing"
)

func TestGoGet(t *testing.T) {
	// Test case 1: Package installation succeeds
	pkg := DependencyDescriptor{
		Name:   "honnef.co/go/tools/cmd/staticcheck",
		Detail: "honnef.co/go/tools/cmd/staticcheck",
	}

	_, err := goGet(pkg)

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
