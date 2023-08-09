package crsce

import crypto "github.com/sam-caldwell/go/v2/projects/crypto/sha256stream"

// CrossSum - A sum/hash pair
type CrossSum struct {
	sum  uint64
	hash *crypto.Sha256Stream
}
