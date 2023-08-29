//go:build arm64
// +build arm64

package hijack

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/convert"
)

// Assembles a jump to a function value
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
	 * What this should create:
	 *
	 *     ldr x17, [destination]
	 *     br x17
	 */
	/*
	 * instructionSize is 16 bytes (128 bits) constant...efficient, eh?
	 * we calculate that as...
	 *
	 *  	'ldr x17' instruction + 4 bytes (32 bits from 0-3)
	 *      <address> bytes       + 8 bytes (64 bits from 4-11) a constant in ARM64
	 *      'br x17'              + 4 bytes (32 bits from 12-15)
	 *                          -----------
	 *                             16 bytes total
	 *
	 * A note on endian-ness...  Arm64 only supports littleEndian.
	 */
	const instructionSize = 16

	instructions := make([]byte, instructionSize)

	// ldr x17, [destination]
	instructions[0] = 0xF8
	instructions[1] = 0x1F
	instructions[2] = 0xFF
	instructions[3] = 0x10

	copy(instructions[4:], convert.UintPtrToByteSlice(destination))

	// br x17
	instructions[12] = 0xE0
	instructions[13] = 0x03
	instructions[14] = 0x3F
	instructions[15] = 0xD6

	return instructions
}
