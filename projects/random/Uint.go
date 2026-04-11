package random

import (
	"math/rand"
)

// Uint generates a random unsigned integer between min (inclusive) and max (exclusive).
func Uint(min, max uint) uint {
	if max <= min {
		return 0
	}

	diff := uint64(max) - uint64(min)
	return uint(rand.Uint64()%diff + uint64(min))
}
