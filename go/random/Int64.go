package random

import (
	"math/rand"
)

// Int64 generates a random int64 between min (inclusive) and max (exclusive).
func Int64(min, max int64) int64 {
	if max <= min {
		return 0
	}
	diff := uint64(max - min)
	return int64(rand.Uint64()%diff + uint64(min))
}
