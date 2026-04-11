package random

import (
	"math"
	"testing"
)

func TestRandomInt(t *testing.T) {
	t.Run("happy path: test multiples", func(t *testing.T) {
		minInput, maxInput := int(-10), int(100)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Int(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Int64(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})
	t.Run("happy path: test range", func(t *testing.T) {
		minInput, maxInput := int(math.MinInt), int(math.MaxInt)
		numSamples := 1000
		for i := 0; i < numSamples; i++ {
			random := Int(minInput, maxInput)
			if random < minInput || random >= maxInput {
				t.Fatalf("Int(%d, %d) = %d; want between %d and %d", minInput, maxInput, random, minInput, maxInput)
			}
		}
	})

	t.Run("Test when max is less than or equal to min", func(t *testing.T) {
		minInput, maxInput := int(math.MaxInt), int(math.MinInt)
		random := Int(minInput, maxInput)
		if random != 0 {
			t.Fatalf("Int(%d, %d) = %d; want 0 as max is less than or equal to min", minInput, maxInput, random)
		}
	})
}
