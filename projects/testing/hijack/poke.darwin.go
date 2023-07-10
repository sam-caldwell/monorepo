//go:build darwin
// +build darwin

package hijack

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"reflect"
	"syscall"
	"unsafe"
)

// poke - Write data (byte slice) to memory at location (HORRIBLY UNSAFE). (Darwin/MacOS Version)
func poke(memoryAddress uintptr, sourceData []byte) (err error) {
	/*
	 * Warning:
	 * This is about as unsafe as you can get.  We're playing with memory here.
	 * 1. We have to disable memory protection first.  This may fail on some systems.
	 * 2. Then we will copy a byte slice to a memory location directly.
	 */
	/*
	 * Capture a reference to our memory space using peek().  Just imagine a cave man in memory.
	 * He saw memory he wanted.  He said "Mine!" and the memory is now his...or we crashed the
	 * application because it turned out the memory he wanted was a pissed off tiger.
	 */
	size := len(sourceData)
	fmt.Printf("poke() calling peek() with ptr:%v [sz:%v]\n", memoryAddress, size)
	//destinationMemory := peek(memoryAddress, size)
	destinationMemory := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: memoryAddress,
		Len:  size,
		Cap:  size,
	}))
	/*
	 * Now comes the fun part...
	 * 1. We will turn off the memory protections (probably piss off endpoint protection) and maybe
	 *    drill a hole in the ground, crashing the app.
	 * 2. Then we will copy our 'sourceData' byte slice to our destinationMemory area.
	 * 3. If we haven't crashed shit by now, we will reset our memory protections to allow READ and EXECUTE only.
	 */
	// make memory read,write,executable
	err = changeMemoryProtectionFlags(memoryAddress, size, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	if err != nil {
		return err
	}
	// copy memory
	fmt.Printf("poke(): destinationMemory: [sz:%0d]: %0x\n", len(destinationMemory), destinationMemory)
	fmt.Printf("poke(): SourceData: %s\n", convert.ByteToHexString(sourceData))

	copy(destinationMemory, sourceData[:])
	// Note: do not reset permissions in darwin...bad things will happen
	//return changeMemoryProtectionFlags(memoryAddress, size, syscall.PROT_READ|syscall.PROT_EXEC)
	return nil
}
