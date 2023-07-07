//go:build arm64
// +build arm64

package hijack

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
	instructions := make([]byte, 12)

	// ldr x17, [destination]
	instructions[0] = 0xF8
	binary.LittleEndian.PutUint64(instructions[1:], uint64(destination))

	// br x17
	copy(instructions[9:], []byte{0xE0, 0x03, 0x3F, 0xD6})

	return instructions
}
