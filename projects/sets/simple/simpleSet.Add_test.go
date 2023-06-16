package simple

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	// Create a new set
	set := Set{}

	// Add the first item (type int)
	err := set.Add(42)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify that the set's type is int
	expectedType := reflect.TypeOf(42)
	if set.typ != expectedType {
		t.Errorf("Expected type %v, got: %v", expectedType, set.typ)
	}

	// Add a second item (type string), which should trigger an error
	err = set.Add("hello")
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	// Verify that the set's type remains int
	if set.typ != expectedType {
		t.Errorf("Expected type %v, got: %v", expectedType, set.typ)
	}

	// Add a third item (type int), which should be added successfully
	err = set.Add(24)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify that the set contains the added items
	expectedData := map[interface{}]bool{
		42: true,
		24: true,
	}
	if !reflect.DeepEqual(set.data, expectedData) {
		t.Errorf("Expected data %v, got: %v", expectedData, set.data)
	}
}
