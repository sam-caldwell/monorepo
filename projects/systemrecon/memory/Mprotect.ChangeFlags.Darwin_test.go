//go:build darwin
// +build darwin

package systemrecon

import (
	"syscall"
	"testing"
	"unsafe"
)

func TestChangeMemoryProtectionFlags(t *testing.T) {
	// Test case: Change memory protection flags
	// Create an object
	testObject := []byte("This is a test object")
	// Get the address of the object and its length
	memoryAddress := uintptr(unsafe.Pointer(&testObject))
	length := len(testObject)
	// try setting our flags
	memoryProtectionFlags := syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC
	err := ChangeFlags(memoryAddress, length, memoryProtectionFlags)
	// Verify that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case: Change memory protection flags for invalid memory address
	invalidMemoryAddress := uintptr(0x0)
	err = ChangeFlags(invalidMemoryAddress, length, memoryProtectionFlags)
	// Verify that an error occurred
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
