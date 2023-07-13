package keyvalue

import (
	"reflect"
	"testing"
)

func TestToOrderedList(t *testing.T) {
	kv := &KeyValue{
		data: Map{
			"foo": 42,
			"bar": "hello",
			"baz": true,
		},
	}

	// Test case 1: Unsorted list
	unsortedResult := kv.ToOrderedList(false)
	expectedUnsorted := []OrderedPair{
		{Key: "foo", Value: 42},
		{Key: "bar", Value: "hello"},
		{Key: "baz", Value: true},
	}
	if !reflect.DeepEqual(unsortedResult, expectedUnsorted) {
		t.Errorf("Unsorted list does not match expected result")
	}

	// Test case 2: Sorted list
	sortedResult := kv.ToOrderedList(true)
	expectedSorted := []OrderedPair{
		{Key: "bar", Value: "hello"},
		{Key: "baz", Value: true},
		{Key: "foo", Value: 42},
	}
	if !reflect.DeepEqual(sortedResult, expectedSorted) {
		t.Errorf("Sorted list does not match expected result")
	}

	// Test case 3: Empty list
	emptyKV := &KeyValue{}
	emptyResult := emptyKV.ToOrderedList(false)
	if len(emptyResult) != 0 {
		t.Errorf("Empty list should have no elements")
	}
}
