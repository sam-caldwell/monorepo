package systemrecon

import (
	"os"
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
