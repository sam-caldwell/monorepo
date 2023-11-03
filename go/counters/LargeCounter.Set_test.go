package counters

/*
 * LargeCounter.Set() Tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the Set() method.
 */

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"math"
	"testing"
)

func TestLargeCounter_Set(t *testing.T) {
	func() {
		//Create a 10-element counter
		c, err := NewLargeCounter(10 * 64)
		if err != nil {
			t.Fatal(err)
		}
		if len(*c) != 10 {
			t.Fatalf("expected size mismatch: %d", len(*c))
		}
		if err := c.Set(0, 1); err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 1 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(0, 2); err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 2 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(0, 3); err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 3 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(0, 4); err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 4 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(1, 1); err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 4 {
			t.Fatal("expected value mismatch")
		}
		if (*c)[1] != 1 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(9, math.MaxUint64); err != nil {
			t.Fatal(err)
		}
		if (*c)[9] != math.MaxUint64 {
			t.Fatal("expected value mismatch")
		}
		if err := c.Set(10, 1); err == nil {
			t.Fatal("expected err not found")
		} else {
			if err.Error() != errors.OverflowError {
				t.Fatal("error mismatch")
			}
		}
	}()
}
