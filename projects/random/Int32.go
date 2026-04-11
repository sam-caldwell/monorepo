package random

import (
	"math/rand"
)

// Int32 generates a random int32 between min (inclusive) and max (exclusive).
func Int32(min, max int32) int32 {
	if max <= min {
		return 0
	}
	diff := uint64(max - min)
	return int32(rand.Uint64()%diff + uint64(min))
}
