package descriptor

import (
	"testing"
)

func TestDescriptor_storeHelpString(t *testing.T) {
	var descriptor Descriptor

	actual := "test help"
	expected := "test help"

	if err := descriptor.storeHelpString(&actual); err != nil {
		t.Fatal(err)
	}
	if result := descriptor.GetHelp(); result != expected {
		t.Fatal("actual/expected mismatch")
	}
}
