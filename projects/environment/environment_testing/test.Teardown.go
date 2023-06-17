package environment_testing

import (
	"os"
	"testing"
)

// TearDown - unset the environment variable
func TearDown(t *testing.T, n string) {
	if err := os.Unsetenv(n); err != nil {
		t.Fatal(err)
	}
}
