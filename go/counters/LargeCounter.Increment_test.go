package counters

/*
 * LargeCounter.Increment() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the increment() method
 */

import (
	"math"
	"testing"
)

func TestLargeCounter_Increment(t *testing.T) {

	func() {
		//Create 64-bit counter
		c, err := NewLargeCounter(64)
		if err != nil {
			t.Fatal(err)
		}
		if len(*c) != 1 {
			t.Fatal("size expects 1 element")
		}
		if (*c)[0] != 0 {
			t.Fatal("initial value expects 0")
		}
		c.Increment()
		if (*c)[0] != 1 {
			t.Fatal("initial value expects 1")
		}
		// Expect rollover like an odometer when we reach max
		(*c)[0] = math.MaxUint64
		c.Increment()
		if (*c)[0] != 0 {
			t.Fatal("initial value expects 0")
		}
	}()

	func() {
		//Create 128-bit counter
		c, err := NewLargeCounter(128)
		if err != nil {
			t.Fatal(err)
		}
		if len(*c) != 2 {
			t.Fatal("size expects 2 element")
		}
		if (*c)[0] != 0 {
			t.Fatal("initial value expects 0")
		}
		c.Increment()
		if (*c)[0] != 1 {
			t.Fatal("initial value expects 1")
		}
		// Expect rollover like an odometer when we reach max
		(*c)[0] = math.MaxUint64
		c.Increment()
		if (*c)[0] != 0 {
			t.Fatal("initial value expects element 0 to be 0")
		}
		if (*c)[1] != 1 {
			t.Fatal("initial value expects element 1 to be 1")
		}
	}()
}
