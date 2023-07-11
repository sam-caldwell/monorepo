package systemrecon

import (
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"testing"
)

func TestGetCurrentHomeDir(t *testing.T) {
	expectedHomeDir := os.Getenv("HOME")
	actualHome, err := GetCurrentHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	if actualHome != expectedHomeDir {
		t.Fatalf("actual not expected home directory.\n"+
			"Expected: %s\n"+
			"Actual: %s", expectedHomeDir, actualHome)
	}
}
