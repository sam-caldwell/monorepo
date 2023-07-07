//go:build arm64
// +build arm64

package hijack

import (
	"reflect"
	"testing"
)

func TestAssemblyJmpToFunction(t *testing.T) {
	// Test case: Destination is a valid address
	destination := uintptr(0x12345678)
	expectedInstructions := []byte{0xF8, 0x1F, 0xFF, 0x10, 0x78, 0x56, 0x34, 0x12, 0xE0, 0x03, 0x3F, 0xD6}
	instructions := assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	if !reflect.DeepEqual(instructions, expectedInstructions) {
		t.Errorf("Generated instructions do not match expected instructions")
	}

	// Test case: Destination is 0 (null address)
	destination = uintptr(0)
	expectedInstructions = []byte{0xF8, 0x1F, 0xFF, 0x10, 0x00, 0x00, 0x00, 0x00, 0xE0, 0x03, 0x3F, 0xD6}
	instructions = assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	if !reflect.DeepEqual(instructions, expectedInstructions) {
		t.Errorf("Generated instructions do not match expected instructions")
	}
}
