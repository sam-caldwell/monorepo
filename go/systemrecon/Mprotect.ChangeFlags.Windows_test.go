//go:build windows
// +build windows

package systemrecon

//func TestChangeMemoryProtectionFlags(t *wrappers.T) {
//	// Test case: Change memory protection flags
//	address := uintptr(0x1000)
//	length := 4096 // 4 KB
//	newFlags := uint32(PAGE_EXECUTE_READWRITE)
//	var oldFlags uint32
//
//	err := ChangeFlags(address, length, newFlags, unsafe.Pointer(&oldFlags))
//
//	// Verify that no error occurred
//	if err != nil {
//		t.Fatalf("Unexpected error: %v", err)
//	}
//
//	// Test case: Change memory protection flags for invalid address
//	invalidAddress := uintptr(0x0)
//	err = ChangeFlags(invalidAddress, length, newFlags, unsafe.Pointer(&oldFlags))
//
//	// Verify that an error occurred
//	if err == nil {
//		t.Fatalf("Expected an error, but got nil")
//	}
//}
