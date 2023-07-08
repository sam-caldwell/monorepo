//go:build amd64
// +build amd64

package hijack

import (
	"unsafe"
)

// assemblyJmpToFunction - Create a byte slice containing our Assembly Language to Jump to a function entrypoint.
func assemblyJmpToFunction(destination uintptr) []byte {
	/*
	 * Theory:
	 * If we know the entrypoint of a function, then we can jump to that
	 * address and execute our new imposter function.  This function aims
	 * to create our assembly instructions to do this and return the same
	 * as a byte slice.
	 *
	 * This could also just mess things up.  But we shall listen to the wise
	 * words of the Van Halen brothers who tell us "You might as well JMP."
	 *
	 * Documentation: https://www.amd.com/system/files/TechDocs/40332.pdf
	 *
	 * What this should create:
	 *
	 *     mov dx, destination
	 *     jmp DWORD PTR [edx]
	 */

	const MOV = 0xBA // Assembly instruction MOV
	const JMP = 0xFF // Assembly instruction JMP
	const RDX = 0x22 // CPU RDX Register

	/*
	 * In golang uintptr could vary in size depending on the system.  It could be a 32-bit
	 * system with 32-bit addresses or a 64-bit system with 64-bit addresses.  We cannot
	 * make an assumption here.  We need to determine the appropriate size then both
	 * create our final byte slice and encode 'destination' (uintptr) into that resulting
	 * byte slice.
	 *
	 * How big is destination?  ("Size does matter...ask a programmer" --Monica <redacted>)
	 */

	size := unsafe.Sizeof(destination)

	/*
	 * Let's go grab our section of memory like a California gold miner in 1849...
	 */

	destBytes := peek(destination, int(size))

	/*
	 * Then let's allocate memory for our instructions.  It needs to be three bytes longer than the size of uintptr.
	 * Why three bytes? Let's visualize it...
	 *
	 *  | pos | contents
	 *  |-----|------------------ - - - -
	 *  |  0  | MOV instruction
	 *  | 1-n | Destination value (uintptr) in []byte
	 *  | n+1 | JMP instruction
	 *  | n+2 | RDX register
	 */

	instructions := make([]byte, size+3)

	/*
	 * Our first Assembly instruction (MOV) starts it all off
	 */

	instructions[0] = MOV

	/*
	 * Next we copy in our slice of bytes from the destination value starting at instruction[1] since we already
	 * put something (MOV) in instruction[0].
	 *
	 * When copy() is done, we will have 'MOV [destination]' in memory
	 */
	//
	copy(instructions[1:], destBytes)

	/*
	 * Now we add 'JMP RDX' to the end of our byte slice and return the result.
	 */

	instructions[size+1] = JMP
	instructions[size+2] = RDX

	return instructions
}
