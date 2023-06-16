package environment

import (
	"os"
	"testing"
)

// setup - test function that creates environment variable (for testing only)
func setup(t *testing.T, n string, v string) {
	if err := os.Setenv(n, v); err != nil {
		t.Fatal(err)
	}
}
