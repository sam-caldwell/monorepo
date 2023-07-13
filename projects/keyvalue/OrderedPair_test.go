package keyvalue

import (
	"reflect"
	"testing"
)

func TestOrderedPair(t *testing.T) {
	pair := OrderedPair{}

	// Test case 1: Verify default values
	expectedKey := ""
	expectedValue := any(nil) // Adjust this with the expected default value type
	if pair.Key != expectedKey {
		t.Errorf("Expected Key to be %s, but got %s", expectedKey, pair.Key)
	}
	if pair.Value != expectedValue {
		t.Errorf("Expected Value to be %v, but got %v", expectedValue, pair.Value)
	}

	// Test case 2: Verify structure
	expectedType := reflect.TypeOf(OrderedPair{})
	if reflect.TypeOf(pair) != expectedType {
		t.Errorf("Expected OrderedPair type to be %s, but got %s", expectedType, reflect.TypeOf(pair))
	}
}
