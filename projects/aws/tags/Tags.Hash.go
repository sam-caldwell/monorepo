package tags

import (
	"github.com/sam-caldwell/monorepo/projects/misc/words"
	"github.com/sam-caldwell/monorepo/projects/types/hashes"
)

// Hash - Return the Sha256 hash of the current tag set
//
//	(c) 2023 Sam Caldwell.  MIT License
func (tags *Tags) Hash() []byte {
	var h hashes.Sha256
	h.HashString(tags.String(words.Colon, words.Comma))
	return h.ToSlice()
}
