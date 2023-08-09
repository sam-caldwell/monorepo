package crsce

import crypto "github.com/sam-caldwell/go/v2/projects/crypto/sha256stream"

// NewCrossSum - Create/initialize a Cross Sum.
func NewCrossSum() *CrossSum {
	var cs CrossSum
	cs.sum = uint64(0)
	cs.hash = crypto.NewSha256Stream()
	return &cs
}
