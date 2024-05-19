//go:build darwin
// +build darwin

package systemrecon

import (
	"testing"
)

func TestChangeMemoryProtectionFlags(t *testing.T) {
	t.Skip("disabled. memoryProtection flags have issues with MacOS security features")
	//// Test case: Change memory protection flags
	//// Create an object
	//testObject := []byte("This is a test object")
	//// Get the address of the object and its length
	//memoryAddress := uintptr(unsafe.Pointer(&testObject))
	//length := len(testObject)
	//// try setting our flags
	//memoryProtectionFlags := syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC
	//err := ChangeFlags(memoryAddress, length, memoryProtectionFlags)
	//// Verify that no error occurred
	//if err != nil {
	//	t.Fatalf("Unexpected error: %v", err)
	//}
	//
	//// Test case: Change memory protection flags for invalid memory address
	//invalidMemoryAddress := uintptr(0x0)
	//err = ChangeFlags(invalidMemoryAddress, length, memoryProtectionFlags)
	//// Verify that an error occurred
	//if err == nil {
	//	t.Fatalf("Expected an error, but got nil")
	//}
}
