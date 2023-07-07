//go:build darwin && linux && windows
// +build darwin,linux,windows

package hijack

import (
	"testing"
)

func TestPoke(t *testing.T) {
	// Test case: Write data to memory
	memoryAddress := uintptr(0x1000)
	sourceData := []byte{0x12, 0x34, 0x56, 0x78}
	err := poke(memoryAddress, sourceData)

	// Verify that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case: Write empty data to memory
	emptyMemoryAddress := uintptr(0x2000)
	emptySourceData := []byte{}
	err = poke(emptyMemoryAddress, emptySourceData)

	// Verify that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
