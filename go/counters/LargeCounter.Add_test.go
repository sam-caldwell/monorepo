package counters

/*
 * LargeCounter.Add() Tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the Add() method
 */

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"math"
	"testing"
)

func TestLargeCounter_Add(t *testing.T) {
	func() {
		//Create 64-bit counter
		c, err := NewLargeCounter(64)
		if err != nil {
			t.Fatal(err)
		}
		if (*c)[0] != 0 {
			t.Fatal("expect initial zero value")
		}
		_ = c.Add(1)
		if (*c)[0] != 1 {
			t.Fatal("expect value 1")
		}
		err = c.Add(math.MaxUint64)
		if err == nil {
			t.Fatal("expected overflow error")
		}
		if err.Error() != errors.OverflowError {
			t.Fatal("error mismatch")
		}
	}()

	func() {
		//10-element 64-bit counter
		c, err := NewLargeCounter(64 * 10)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < len(*c); i++ {
			if (*c)[i] != 0 {
				t.Fatalf("expect 0-value for element %d (got: %v)", i, (*c)[i])
			}
		}
		err = c.Add(math.MaxUint64)
		if (*c)[0] != math.MaxUint64 {
			t.Fatal("expected element-0 value mismatch")
		}
		for i := 1; i < len(*c); i++ {
			if (*c)[i] != 0 {
				t.Fatalf("expect 0-value for element %d (got: %v)", i, (*c)[i])
			}
		}
		err = c.Add(2)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if (*c)[0] != 1 {
			t.Fatal("expected element-0 value mismatch")
		}
		if (*c)[1] != 1 {
			t.Fatal("expected element-1 value mismatch")
		}
		//Max all elements
		for i := 0; i < len(*c); i++ {
			(*c)[i] = math.MaxUint64
		}
		err = c.Add(1)
		if err == nil {
			t.Fatal("expect overflow error")
		}
		if err.Error() != errors.OverflowError {
			t.Fatal("overflow error message mismatch")
		}
	}()
}
