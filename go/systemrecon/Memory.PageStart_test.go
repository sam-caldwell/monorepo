package systemrecon

import (
	"testing"
)

func TestPageStart(t *testing.T) {
	// Test case: Pointer within a page
	ptr := uintptr(0x12345678)
	expectedStart := uintptr(0x12345000) // Assuming page size is 4096 bytes
	start := PageStart(ptr)

	// Verify that the calculated start address matches the expected start address
	if start != expectedStart {
		t.Fatalf("Expected start address: %#x, but got: %#x", expectedStart, start)
	}

	// Test case: Pointer at the start of a page
	ptr = uintptr(0x80000000)
	expectedStart = uintptr(0x80000000) // Assuming page size is 4096 bytes
	start = PageStart(ptr)

	// Verify that the calculated start address matches the expected start address
	if start != expectedStart {
		t.Fatalf("Expected start address: %#x, but got: %#x", expectedStart, start)
	}

	// Test case: Pointer at the end of a page
	ptr = uintptr(0xABCDFFFF)
	expectedStart = uintptr(0xABCDF000) // Assuming page size is 4096 bytes
	start = PageStart(ptr)

	// Verify that the calculated start address matches the expected start address
	if start != expectedStart {
		t.Fatalf("Expected start address: %#x, but got: %#x", expectedStart, start)
	}
}
