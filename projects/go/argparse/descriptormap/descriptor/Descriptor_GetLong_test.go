package descriptor

import (
	"testing"
)

// TestDescriptor_GetLong - Test GetLong() and assume no sanitization when directly setting param
func TestDescriptor_GetLong(t *testing.T) {
	test := func(input string) {
		var descriptor Descriptor
		descriptor.long = input
		if result := descriptor.GetLong(); result != input {
			t.Fatalf("GetLong() mismatch: %s %s", result, input)
		}
	}
	test("--long-arg")
	test("--long")
	test("--long_arg")
	test("returns anything on GetLong()")
}
