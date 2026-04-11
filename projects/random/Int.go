package random

import (
	"math/rand"
)

// Int generates a random integer between min (inclusive) and max (exclusive).
func Int(min, max int) int {
	if max <= min {
		return 0
	}
	diff := uint64(max - min)
	return int(rand.Uint64()%diff + uint64(min))
}
