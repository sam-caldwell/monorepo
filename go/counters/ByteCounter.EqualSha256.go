package counters

import (
	"bytes"
	"crypto/sha256"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
)

// EqualSha256 - Compare two ByteCounters using a sha256 hash of their content
func (c *ByteCounter) EqualSha256(rhs *ByteCounter) bool {

	var a hashes.Sha256
	var b hashes.Sha256

	c.lock.Lock()
	rhs.lock.Lock()

	defer c.lock.Unlock()
	defer rhs.lock.Unlock()

	a = sha256.Sum256(c.v)
	b = sha256.Sum256(rhs.v)
	return bytes.Equal(a[:], b[:])

}
