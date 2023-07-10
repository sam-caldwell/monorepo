package systemrecon

/*
 * projects/systemrecon/memory/Peek.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Peek() will capture a given number (length) of bytes starting at
 * a specified address (p) and return the result as a []byte slice.
 * Note, we do this in a way that we return the actual section of
 * memory.
 */

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Peek - It's raw memory access golang style
func Peek(p uintptr, length int) *[]byte {
	fmt.Print("peek(): ", p, length, "\n")
	/*
	 * Remember GW-BASIC?  No.  Oh, you're too young then
	 * BASIC had a command called peek() and poke() which
	 * where my first forays into memory-level fun.
	 */
	/*
	 * What this does...
	 * 1. We create a slice using reflect.SliceHeader of 'length' bytes at memory location p.
	 *    Imagine that it just drew a box around a bunch of memory and said "MINE!" because that's what it did.
	 * 2. Then it uses unsafe.Pointer() to convert the newly created slice into a []byte slice, providing us
	 *    with direct access to that memory.
	 * 3. finally, it returns a pointer to this newly defined object.
	 */
	memory := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
	return &memory
	/*
	 * Note: This is YOLO land.  There is no error handling because at this level, you're in the wild west
	 *       and any errors are just BOHICA time.
	 */
}
