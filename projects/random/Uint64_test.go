package random

import (
	"math"
	"testing"
)

func TestRandomUint64(t *testing.T) {
	t.Run("happy path: test multiples", func(t *testing.T) {
		minInput, maxInput := uint64(0), uint64(100)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint64(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Int64(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})
	t.Run("happy path: test range", func(t *testing.T) {
		minInput, maxInput := uint64(0), uint64(math.MaxUint64)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint64(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Uint64(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})

	t.Run("Test when max is less than or equal to min", func(t *testing.T) {
		minInput, maxInput := uint64(math.MaxUint64), uint64(1)
		random := Uint64(minInput, maxInput)
		if random != 0 {
			t.Fatalf("Uint64(%d, %d) = %d; want 0 as max is less than or equal to min", minInput, maxInput, random)
		}
	})
}
