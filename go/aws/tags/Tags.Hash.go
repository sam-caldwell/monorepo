package tags

import "github.com/sam-caldwell/monorepo/go/types/hashes"

// Hash - Return the Sha256 hash of the current tag set
func (tags *Tags) Hash() []byte {
	var h hashes.Sha256
	h.HashString(tags.String())
	return h.ToSlice()
}
