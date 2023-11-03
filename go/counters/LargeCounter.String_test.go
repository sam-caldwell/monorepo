package counters

/*
 * LargeCounter.String() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the String() method
 */

import "testing"

func TestLargeCounterString(t *testing.T) {
	// Create a LargeCounter with some sample data
	counter := LargeCounter{123, 456, 789}

	// Calculate the expected hexadecimal string
	expected := "000000000000007b00000000000001c80000000000000315"

	// Get the actual hexadecimal string
	actual := counter.String()

	// Check if the actual result matches the expected result
	if actual != expected {
		t.Errorf("String() returned %s, expected %s", actual, expected)
	}
}
