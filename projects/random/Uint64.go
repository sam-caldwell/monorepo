package random

import (
	"math"
	"math/rand"
)

// Uint64 generates a random unsigned 64-bit between min (inclusive) and max (exclusive).
func Uint64(min, max uint64) uint64 {
	if max <= min {
		return 0
	}

	diff := max - min
	if diff > math.MaxInt64 {
		return min + rand.Uint64()%diff
	}

	return min + uint64(rand.Int63n(int64(diff)))
}
