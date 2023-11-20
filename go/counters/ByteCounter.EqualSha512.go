package counters

import (
	"bytes"
	"crypto/sha512"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
)

// EqualSha512 - Compare two ByteCounters using a Sha512 hash of their content
func (c *ByteCounter) EqualSha512(rhs *ByteCounter) bool {

	var a hashes.Sha512
	var b hashes.Sha512

	c.lock.Lock()
	rhs.lock.Lock()

	defer c.lock.Unlock()
	defer rhs.lock.Unlock()

	a = sha512.Sum512(c.v)
	b = sha512.Sum512(rhs.v)
	return bytes.Equal(a[:], b[:])

}
