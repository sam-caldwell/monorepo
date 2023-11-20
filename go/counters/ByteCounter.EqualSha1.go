package counters

import (
	"bytes"
	"crypto/sha1"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
)

// EqualSha1 - Compare two ByteCounters using a sha1 hash of their content
func (c *ByteCounter) EqualSha1(rhs *ByteCounter) bool {

	var a hashes.Sha1
	var b hashes.Sha1

	c.lock.Lock()
	rhs.lock.Lock()

	defer c.lock.Unlock()
	defer rhs.lock.Unlock()

	a = sha1.Sum(c.v)
	b = sha1.Sum(rhs.v)
	return bytes.Equal(a[:], b[:])

}
