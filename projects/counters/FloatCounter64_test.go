package counters

import "testing"

func TestFloatCounter64(t *testing.T) {
	increments := []float64{-1.0, -0.5, -0.25, -0.1, -0.01, 0.01, 0.1, 0.25, 0.5, 1.0}

	for _, increment := range increments {
		counter := FloatCounter64(0.0, increment)

		// Perform five increments
		for i := 1; i <= 5; i++ {
			expected := float64(i) * increment
			result := counter()

			if result != expected {
				t.Errorf("Counter value after %d increments with increment %f is %f, expected %f",
					i, increment, result, expected)
			}
		}
	}
}
