package environment_testing

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"testing"
)

// Setup - test function that creates environment variable (for testing only)
func Setup(t *testing.T, n string, v string) {
	if err := os.Setenv(n, v); err != nil {
		t.Fatal(err)
	}
}
