//go:build linux
// +build linux

package hijack

import (
	"syscall"
	"testing"
)

func TestChangeMemoryProtectionFlags(t *testing.T) {
	// Test case: Change memory protection flags
	memoryAddress := uintptr(0x1000)
	length := 4096 // 4 KB
	memoryProtectionFlags := syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC

	err := changeMemoryProtectionFlags(memoryAddress, length, memoryProtectionFlags)

	// Verify that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case: Change memory protection flags for invalid memory address
	invalidMemoryAddress := uintptr(0x0)
	err = changeMemoryProtectionFlags(invalidMemoryAddress, length, memoryProtectionFlags)

	// Verify that an error occurred
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
