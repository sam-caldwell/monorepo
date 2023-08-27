package moremath

import (
	"testing"
)

func TestTruncFloat32(t *testing.T) {
	// Test cases with positive values
	value := float32(3.14159)
	position := 2
	expected := float32(3.14)
	result := TruncFloat32(value, position)
	if result != expected {
		t.Errorf("TruncFloat32(%f, %d) = %f, expected %f", value, position, result, expected)
	}

	// Test cases with negative values
	value = float32(-5.678)
	position = 1
	expected = float32(-5.6)
	result = TruncFloat32(value, position)
	if result != expected {
		t.Errorf("TruncFloat32(%f, %d) = %f, expected %f", value, position, result, expected)
	}

	// Test cases with zero value
	value = float32(0)
	position = 3
	expected = float32(0)
	result = TruncFloat32(value, position)
	if result != expected {
		t.Errorf("TruncFloat32(%f, %d) = %f, expected %f", value, position, result, expected)
	}
}
