package systemrecon

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"testing"
)

func TestGetCurrentUserName(t *testing.T) {
	expected := os.Getenv("USER")
	actual, err := GetCurrentUserName()
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("actual not expected user.\n"+
			"Expected: %s\n"+
			"Actual: %s", expected, actual)
	}
}
