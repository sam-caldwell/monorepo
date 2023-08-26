package descriptor

import (
	"strings"
	"testing"
)

// TestStoreShort - test Descriptor.storeShort()
func TestStoreShort(t *testing.T) {
	test := func(arg string, expectedError bool) {
		var descriptor Descriptor
		if err := descriptor.storeShort(&arg); err != nil {
			if !expectedError {
				t.Fatal(err)
			}
		}
		expected := strings.ToLower(strings.TrimSpace(strings.TrimSpace(arg)))

		if arg != expected {
			t.Fatalf("short mismatch (%s)(%s)", arg, expected)
		}

	}
	test("", false)
	test("-a", false)
	test("--arg", true)
	test("invalid", true)

}
