package counters

import (
	"reflect"
	"testing"
)

func TestLargeCounter_Bytes(t *testing.T) {
	// Create a LargeCounter with some sample data
	counter := LargeCounter{123, 456, 789}

	// Calculate the expected byte representation in big-endian order
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 123, // 123 in big-endian
		0, 0, 0, 0, 0, 0, 1, 200, // 456 in big-endian
		0, 0, 0, 0, 0, 0, 3, 21} // 789 in big-endian

	// Get the actual byte representation
	actual := counter.Bytes()

	// Check if the actual result matches the expected result
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Bytes() returned %v, expected %v", actual, expected)
	}
}
