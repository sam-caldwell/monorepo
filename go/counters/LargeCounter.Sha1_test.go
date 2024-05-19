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
	expected := "c769d306d87e750b9d5358b8dc347ce59c12b1fb"

	// Get the actual SHA-1 hash
	actual := counter.Sha1()

	// Check if the actual result matches the expected result
	if actual != expected {
		t.Fatalf("Sha1() returned\n"+
			"actual:   %s,\n"+
			"expected: %s",
			actual, expected)
	}
}
