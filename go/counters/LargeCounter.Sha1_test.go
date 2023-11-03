package counters

/*
 * LargeCounter.Sha1() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the Bytes() method
 */

import (
	"testing"
)

func TestLargeCounter_Sha1(t *testing.T) {
	// Create a LargeCounter with some sample data
	counter := LargeCounter{123, 456, 789}

	// Calculate the expected SHA-1 hash
	expected := "d9f553e5606e5d58adf6f4f01302af12fb6cb8cd"

	// Get the actual SHA-1 hash
	actual := counter.Sha1()

	// Check if the actual result matches the expected result
	if actual != expected {
		t.Errorf("Sha1() returned %s, expected %s", actual, expected)
	}
}
