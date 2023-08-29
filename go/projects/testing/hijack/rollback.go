package hijack

// rollback overwrite the target memory with the original content
func rollback(target uintptr, ap AppliedPatch) (err error) {
	/*
	 * I was feeling nostalgic to my old days playing
	 * with Peek() and Poke() in BASIC.  If you see
	 * this, Dan Duley, you're missed friend.  May you
	 * have continued your quest to write amazing code.
	 */
	//err = systemrecon.Poke(target, ap.originalBytes)
	return err
}
