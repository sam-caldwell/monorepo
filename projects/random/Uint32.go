package random

import (
	"math/rand"
)

// Uint32 generates a random unsigned 32-bit integer between min (inclusive) and max (exclusive).
func Uint32(min, max uint32) uint32 {
	if max <= min {
		return 0
	}
	diff := uint64(max) - uint64(min)
	return uint32(rand.Uint64()%diff + uint64(min))
}
