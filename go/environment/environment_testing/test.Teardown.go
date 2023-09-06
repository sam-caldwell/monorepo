package environment_testing

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

// TearDown - unset the environment variable
func TearDown(t *testing.T, n string) {
	if err := os.Unsetenv(n); err != nil {
		t.Fatal(err)
	}
}
