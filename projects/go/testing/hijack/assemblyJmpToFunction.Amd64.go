//go:build amd64
// +build amd64

package hijack

import (
	"github.com/sam-caldwell/go/v2/projects/go/convert"
)

// AssemblyJmpToFunction - Create a byte slice containing our Assembly Language to Jump to a function entrypoint.
func AssemblyJmpToFunction(destination uintptr) (jumpCode []byte) {
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

	const (
		MOVRDX = 0x48 // Target Register (RDX)
		MOV    = 0xBA // Assembly instruction MOV
		JMP    = 0xFF // Assembly instruction JMP
		RDX    = 0x22 // Destination Register (RDX)
	)

	/*
	 * Convert our destination uintptr address to a []byte slice
	 * then determine the size from the slice length.
	 */

	var destBytes = convert.UintPtrToByteSlice(destination)
	size := len(destBytes)

	/*
	 * Then let's allocate memory for our instructions.  It needs to be three (3) bytes longer than the size
	 * of uintptr.  Why three bytes? Let's visualize it...
	 *
	 *  | pos | contents
	 *  |-----|------------------ - - - -
	 *  |  0  | MOV instruction
	 *  | 1-n | Destination value (uintptr) in []byte
	 *  | n+1 | JMP instruction
	 *  | n+2 | RDX register
	 *
	 * In the above table we see that the size of our destination value plus three instructions at 0,n+1 and n+2.
	 * so the size of destination + 3 is our total instruction[] size.
	 */

	jumpCode = make([]byte, size+4)

	/*
	 * Our first Assembly instruction (MOV) starts it all off
	 */
	jumpCode[0] = MOV
	jumpCode[1] = MOVRDX

	/*
	 * Next we copy in our slice of bytes from the destination value starting at instruction[1] since we already
	 * put something (MOV) in instruction[0].
	 *
	 * When copy() is done, we will have 'MOV [destination]' in memory
	 */
	copy(jumpCode[2:], destBytes)

	/*
	 * Now we add 'JMP RDX' to the end of our byte slice and return the result.
	 */
	jumpCode[size+2] = JMP
	jumpCode[size+3] = RDX

	return jumpCode
}
