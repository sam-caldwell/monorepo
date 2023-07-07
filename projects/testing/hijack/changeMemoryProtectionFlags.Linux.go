//go:build linux
// +build linux

package hijack

import "syscall"

// changeMemoryProtectionFlags - Change our memory protection flags (Darwin/MacOS version)
func changeMemoryProtectionFlags(memoryAddress uintptr, length int, memoryProtectionFlags int) error {
	/*
	 * Starting at 'memoryAddress', traverse the RAM memory for 'length' pages
	 * and edit the memory protection flags for each page using syscall()
	 *
	 * This may fail or piss off your endpoint protection.  It probably should.
	 *
	 * Example flags:
	 * 		syscall.PROT_READ   - read access
	 *  	syscall.PROT_WRITE  - write access
	 *  	syscall.PROT_EXEC   - execute access.
	 *
	 * These flags are passed as `memoryProtectionFlags` (integer) using bitwise OR operations.
	 * For example, to make a page read-only, change protections to PROD_READ, but to make it
	 * read-write, use--
	 *
	 * memoryProtectionFlags:=syscall.PROT_READ | syscall.PROD_WRITE
	 */
	pageSize := syscall.Getpagesize()

	for p := pageStart(memoryAddress); p < memoryAddress+uintptr(length); p += uintptr(pageSize) {
		page := peek(p, pageSize)
		if err := syscall.Mprotect(page, memoryProtectionFlags); err != nil {
			return err
		}
	}
	return nil
}
