//go:build amd64
// +build amd64

package hijack

import (
	"reflect"
	"testing"
)

func TestAssemblyJmpToFunction(t *testing.T) {
	// Test case: Destination is a valid address
	destination := uintptr(0x12345678)
	expectedInstructions := []byte{0xBA, 0x78, 0x56, 0x34, 0x12, 0xFF, 0x22}
	instructions := assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	if !reflect.DeepEqual(instructions, expectedInstructions) {
		t.Errorf("Generated instructions do not match expected instructions")
	}

	// Test case: Destination is 0 (null address)
	destination = uintptr(0)
	expectedInstructions = []byte{0xBA, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x22}
	instructions = assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	if !reflect.DeepEqual(instructions, expectedInstructions) {
		t.Errorf("Generated instructions do not match expected instructions")
	}
}
