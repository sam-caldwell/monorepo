package crypto

import (
	"reflect"
	"testing"
)

func TestSha256StreamOutput(t *testing.T) {
	// Create a test Sha256Stream instance with some arbitrary hash values.
	// Replace these values with the actual test values if you have specific inputs in mind.
	stream := Sha256Stream{
		h: [8]uint32{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19},
	}

	// Calculate the expected output byte slice manually based on the test values above.
	expectedOutput := []byte{
		0x6a, 0x09, 0xe6, 0x67, 0xbb, 0x67, 0xae, 0x85,
		0x3c, 0x6e, 0xf3, 0x72, 0xa5, 0x4f, 0xf5, 0x3a,
		0x51, 0x0e, 0x52, 0x7f, 0x9b, 0x05, 0x68, 0x8c,
		0x1f, 0x83, 0xd9, 0xab, 0x5b, 0xe0, 0xcd, 0x19,
	}

	// Call the Output() method to get the actual output byte slice.
	actualOutput := stream.Output()

	// Compare the actual output with the expected output.
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("Output() method returned incorrect result.\nExpected: %x\nActual: %x", expectedOutput, actualOutput)
	}
}
