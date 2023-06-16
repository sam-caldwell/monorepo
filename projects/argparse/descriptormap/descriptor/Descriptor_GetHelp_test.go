package descriptor

import (
	"testing"
)

func TestDescriptor_GetHelp(t *testing.T) {
	var descriptor Descriptor
	const (
		actual   = "test help"
		expected = "test help"
	)
	descriptor.help = actual
	if result := descriptor.GetHelp(); result != expected {
		t.Fatalf("GetHelp() mismatch: %s %s", result, expected)
	}
}
