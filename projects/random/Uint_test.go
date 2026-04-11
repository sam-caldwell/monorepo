package random

import (
	"math"
	"testing"
)

func TestRandomUint(t *testing.T) {
	t.Run("happy path: test multiples", func(t *testing.T) {
		minInput, maxInput := uint(0), uint(100)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Int64(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})
	t.Run("happy path: test range", func(t *testing.T) {
		minInput, maxInput := uint(0), uint(math.MaxUint)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Uint(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Uint(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})

	t.Run("Test when max is less than or equal to min", func(t *testing.T) {
		minInput, maxInput := uint(math.MaxUint), uint(1)
		random := Uint(minInput, maxInput)
		if random != 0 {
			t.Fatalf("Uint(%d, %d) = %d; want 0 as max is less than or equal to min", minInput, maxInput, random)
		}
	})
}
