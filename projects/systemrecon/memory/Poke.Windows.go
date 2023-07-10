//go:build windows
// +build windows

package systemrecon

import (
	"unsafe"
)

// Poke - Write data (byte slice) to memory at location (HORRIBLY UNSAFE). (Micro$oft Windows Version)
func Poke(memoryAddress uintptr, sourceData []byte) (err error) {
	/*
	 * Warning:
	 * This is about as unsafe as you can get.  We're playing with memory here.
	 * 1. We have to disable memory protection first.  This may fail on some systems.
	 * 2. Then we will copy a byte slice to a memory location directly.
	 */

	//See https://learn.microsoft.com/en-us/windows/win32/memory/memory-protection-constants
	const PageExecuteReadwrite = 0x40 // PAGE_EXECUTE_READWRITE
	var originalPermissions uint32    // This will capture initial state we can restore to when we are done.
	var waste uint32                  //This is a trash var to pass by reference for our Windows API call

	/*
	 * Capture a reference to our memory space using peek().  Just imagine a cave man in memory.
	 * He saw memory he wanted.  He said "Mine!" and the memory is now his...or we crashed the
	 * application because it turned out the memory he wanted was a pissed off tiger.
	 */
	size := len(sourceData)
	destinationMemory := peek(memoryAddress, size)

	// make memory read,write,executable and capture the originalPermissions
	err = ChangeFlags(memoryAddress, size, PageExecuteReadwrite, unsafe.Pointer(&originalPermissions))
	if err != nil {
		return err
	}

	//copy memory
	copy(destinationMemory, sourceData[:])

	// Revert our permissions (Note: we have to pass over an object to capture the permissions)
	return ChangeFlags(memoryAddress, size, originalPermissions, unsafe.Pointer(&waste))
}
