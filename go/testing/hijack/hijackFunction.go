package hijack

// TestHijack - for testing
func TestHijack(targetFuncPtr, imposterFuncPtr uintptr) (originalMemory []byte, err error) {
	return hijackFunction(targetFuncPtr, imposterFuncPtr)
}

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
	//ansi.Blue().Println("hijackFunction() start").
	//	Printf(">>  targetFuncPtr: %0x", targetFuncPtr).LF().
	//	Printf(">>imposterFuncPtr: %0x", imposterFuncPtr).LF().Reset()

	/*
	 * Get the assembly language code we need to jump into our imposter function. This is a jump
	 * code that will be used to redirect execution.
	 */
	//jumpData := AssemblyJmpToFunction(imposterFuncPtr)
	/*
	 * Like a dog marking territory, let's draw a line around the territory we want to claim for
	 * our own.  This is the piece of memory we will be overwriting during our hijack process.
	 */
	//blockSize := len(jumpData)
	//hijackedMemory := systemrecon.Peek(targetFuncPtr, blockSize)
	/*
	 * Now let's copy the mapped section of memory into a []byte slice we can return
	 * to the caller so the change can be undone down the road.
	 */
	//originalMemory = make([]byte, blockSize)
	//copy(originalMemory, *hijackedMemory)

	/*
	 * This is a hijacking, people!  this next step overwrites the first few bytes
	 * of the target function with our assembly language instructions to jump to the
	 * imposter function.  Thank you for flying nerdsurgency airlines, where we will
	 * not only allow you to monkey-patch golang. We will allow you go guerilla-patch
	 * your code.
	 */
	//ansi.Yellow().
	//	LF().
	//	Print("Overwriting target:\n").
	//	Printf("\ttargetFuncPtr:   %0x\n", targetFuncPtr).
	//	Printf("\timposterFuncPtr: %0x\n", imposterFuncPtr).
	//	Printf("\tjumpData [%s]\n", convert.ByteToHexString(jumpData)).
	//	Reset()
	//
	//err = systemrecon.Poke(targetFuncPtr, jumpData)
	/*
	 * Return our error and original memory
	 */
	//ansi.Green().
	//	Printf("returning originalMemory(err:%v) original memory: %s\n", err, convert.ByteToHexString(originalMemory)).
	//	Reset()
	return originalMemory, err
}
