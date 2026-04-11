package random

import (
	"math"
	"testing"
)

func TestRandomUint32(t *testing.T) {
	t.Run("happy path: test multiples", func(t *testing.T) {
		minInput, maxInput := uint32(0), uint32(100)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint32(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Int64(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})
	t.Run("happy path: test range", func(t *testing.T) {
		minInput, maxInput := uint32(0), uint32(math.MaxUint32)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint32(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Uint32(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})

	t.Run("Test when max is less than or equal to min", func(t *testing.T) {
		minInput, maxInput := uint32(math.MaxUint32), uint32(1)
		random := Uint32(minInput, maxInput)
		if random != 0 {
			t.Fatalf("Uint32(%d, %d) = %d; want 0 as max is less than or equal to min", minInput, maxInput, random)
		}
	})
}
