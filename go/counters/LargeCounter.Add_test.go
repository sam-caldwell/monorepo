package counters

/*
 * LargeCounter.Add() Tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the Add() method
 */

import "testing"

func TestLargeCounter_Add(t *testing.T) {
	var o LargeCounter

	if o != nil {
		t.Fatal("Expect nil value initially")
	}
	o = make(LargeCounter, 10)
	if len(o) < 10 {
		t.Fatal("Expect 10-element counter")
	}
}
