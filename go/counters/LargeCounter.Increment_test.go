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
	"time"
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

	func() {
		counter := LargeCounter{0, 0, 0}
		counter.Increment()
		if counter[0] != 1 {
			t.Fatal("incremented wrong position")
		}
	}()

	func() {
		const start = math.MaxUint64
		c, _ := NewLargeCounter(128)
		for i := 0; i < 1023; i++ {
			c.Increment()
		}
		if ((*c)[0] == 1022) && ((*c)[1] == 1) {
			t.Fatal("unexpected outcome")
		}
	}()

	func() {
		const iterations = 10485760
		for pass := 0; pass < 100; pass++ {
			startTime := time.Now().UnixNano()
			c, err := NewLargeCounter(64 * 10)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < iterations; i++ {
				c.Increment()
			}
			if (*c)[0] != iterations {
				t.Fatal("outcome unexpected")
			}
			stopTime := time.Now().UnixNano()
			elapsedPerIteration := float64(stopTime-startTime) / float64(iterations)
			//t.Logf("elapsedPerIteration: %f", elapsedPerIteration)
			if elapsedPerIteration > 10 {
				t.Fatalf("baseline performance not expected (%f) ns/iteration", elapsedPerIteration)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
