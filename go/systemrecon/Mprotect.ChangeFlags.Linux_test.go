//go:build linux
// +build linux

package systemrecon

import (
	"testing"
)

func TestChangeMemoryProtectionFlags(t *testing.T) {
	t.Skip("not implemented")
	//	// Test case: Change memory protection flags
	//	memoryAddress := uintptr(0x1000)
	//	length := 4096 // 4 KB
	//	memoryProtectionFlags := syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC
	//
	//	err := ChangeFlags(memoryAddress, length, memoryProtectionFlags)
	//
	//	// Verify that no error occurred
	//	if err != nil {
	//		t.Fatalf("Unexpected error: %v", err)
	//	}
	//
	//	// Test case: Change memory protection flags for invalid memory address
	//	invalidMemoryAddress := uintptr(0x0)
	//	err = ChangeFlags(invalidMemoryAddress, length, memoryProtectionFlags)
	//
	//	// Verify that an error occurred
	//	if err == nil {
	//		t.Fatalf("Expected an error, but got nil")
	//	}
}
