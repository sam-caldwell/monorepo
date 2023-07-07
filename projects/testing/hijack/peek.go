package hijack

import (
	"reflect"
	"unsafe"
)

// peek - It's raw memory access golang style
func peek(p uintptr, length int) []byte {
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
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
	/*
	 * Note: This is YOLO land.  There is no error handling because at this level, you're in the wild west
	 *       and any errors are just BOHICA time.
	 */
}
