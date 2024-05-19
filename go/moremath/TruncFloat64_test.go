package moremath

import (
	"testing"
)

func TestTruncFloat64(t *testing.T) {
	// Test cases with positive values
	value := 3.14159
	position := 2
	expected := 3.14
	result := TruncFloat64(value, position)
	if result != expected {
		t.Fatalf("TruncFloat64(%f, %d) = %f, expected %f", value, position, result, expected)
	}

	// Test cases with negative values
	value = -5.678
	position = 1
	expected = -5.6
	result = TruncFloat64(value, position)
	if result != expected {
		t.Fatalf("TruncFloat64(%f, %d) = %f, expected %f", value, position, result, expected)
	}

	// Test cases with zero value
	value = 0
	position = 3
	expected = 0
	result = TruncFloat64(value, position)
	if result != expected {
		t.Fatalf("TruncFloat64(%f, %d) = %f, expected %f", value, position, result, expected)
	}
}
