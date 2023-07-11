package hijack

import systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/memory"

// rollback overwrite the target memory with the original content
func rollback(target uintptr, ap AppliedPatch) error {
	/*
	 * I was feeling nostalgic to my old days playing
	 * with Peek() and Poke() in BASIC.  If you see
	 * this, Dan Duley, you're missed friend.  May you
	 * have continued your quest to write amazing code.
	 */
	return systemrecon.Poke(target, ap.originalBytes)
}
