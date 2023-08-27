package descriptor

import (
	"testing"
)

// TestDescriptor_GetShort - Test GetShort
func TestDescriptor_GetShort(t *testing.T) {
	var descriptor Descriptor

	const (
		actual   = "--short_arg--"
		expected = "--short_arg--"
	)

	descriptor.short = actual
	if result := descriptor.GetShort(); result != expected {
		t.Fatalf("GetShort() mismatch: %s %s", result, expected)
	}
}
