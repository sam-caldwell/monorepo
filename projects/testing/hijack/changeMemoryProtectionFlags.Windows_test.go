//go:build windows
// +build windows

package hijack

import (
	"testing"
	"unsafe"
)

func TestChangeMemoryProtectionFlags(t *testing.T) {
	// Test case: Change memory protection flags
	address := uintptr(0x1000)
	length := 4096 // 4 KB
	newFlags := uint32(PAGE_EXECUTE_READWRITE)
	var oldFlags uint32

	err := changeMemoryProtectionFlags(address, length, newFlags, unsafe.Pointer(&oldFlags))

	// Verify that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case: Change memory protection flags for invalid address
	invalidAddress := uintptr(0x0)
	err = changeMemoryProtectionFlags(invalidAddress, length, newFlags, unsafe.Pointer(&oldFlags))

	// Verify that an error occurred
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
