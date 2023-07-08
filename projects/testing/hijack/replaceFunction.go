package hijack

// hijackFunction - overwrite a target function with an imposter go function
func hijackFunction(targetFuncPtr, imposterFuncPtr uintptr) (originalMemory []byte, err error) {
	/*
	 * This function will create a small "jump code" (assembly language) routine (only a few bytes)
	 * then backup the same number of bytes at the start of the target function's head for return
	 * to the caller before overwriting the target function's entrypoint with the jump code.
	 *
	 * There are a few reasons for this approach:
	 *     1. Overwriting the entrypoint with a new entrypoint pointing to the imposter allows
	 *        us to alter functionality as desired ("guerilla patching," folks).
	 *     2. We can restore the old functionality when we are done.
	 *     3. We didn't overwrite the entire target function, nor did we back up the entire
	 *        target function.  This would be potentially expensive.  Instead, our TrackerTable
	 *        records will be a consistent size for less memory impact.
	 */

	/*
	 * Get the assembly language code we need to jump into our imposter function. This is a jump
	 * code that will be used to redirect execution.
	 */
	jumpData := assemblyJmpToFunction(imposterFuncPtr)

	/*
	 * Like a dog marking territory, let's draw a line around the territory we want to claim for
	 * our own.  This is the piece of memory we will be overwriting during our hijack process.
	 */
	hijackedMemory := peek(targetFuncPtr, len(jumpData))
	/*
	 * Now let's copy the mapped section of memory into a []byte slice we can return
	 * to the caller so the change can be undone down the road.
	 */
	originalMemory = make([]byte, len(hijackedMemory))
	copy(originalMemory, hijackedMemory)
	/*
	 * This is a hijacking, people!  this next step overwrites the first few bytes
	 * of the target function with our assembly language instructions to jump to the
	 * imposter function.  Thank you for flying nerdsurgency airlines, where we will
	 * not only allow you to monkey-patch golang. We will allow you go guerilla-patch
	 * your code.
	 */
	err = poke(targetFuncPtr, jumpData)
	/*
	 * Return our error and original memory
	 */
	return originalMemory, err
}
