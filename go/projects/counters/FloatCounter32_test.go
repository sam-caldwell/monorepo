package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/testing/testmessage"
	"testing"
)

func TestFloatCounter32(t *testing.T) {
	test := testmessage.Test(t)
	defer test.Stats()
	increments := []float32{-1.0, -0.5, -0.25, -0.1, -0.01, 0.01, 0.1, 0.25, 0.5, 1.0}

	for _, increment := range increments {
		counter := FloatCounter32(0.0, increment)

		// Perform five increments
		for i := 1; i <= 5; i++ {
			expected := float32(i) * increment
			result := counter()

			if result != expected {
				test.Fatalf("Counter value after %d increments with increment %f is %f, expected %f\n",
					i, increment, result, expected)
			}
			//test.Passf("Counter at %d (%v)\n", i, increment)
		}
	}
}
