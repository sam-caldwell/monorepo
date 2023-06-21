package keyvalue

import (
	"reflect"
	"testing"
)

func TestMergeFromKv(t *testing.T) {
	// Initialize two KeyValue objects with some data
	kv1 := &KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	kv2 := &KeyValue{
		data: Map{
			"key2": "value2new",
			"key3": "value3",
			"key4": "value4",
		},
	}

	// Call the MergeFromKv function
	kv1.MergeFromKv(*kv2)

	// Create a Map with the expected result after the merge
	expected := Map{
		"key1": "value1",
		"key2": "value2new",
		"key3": "value3",
		"key4": "value4",
	}

	// Check if the data in kv1 matches the expected result
	if !reflect.DeepEqual(kv1.data, expected) {
		t.Errorf("MergeFromKv() failed, expected %v, got %v", expected, kv1.data)
	}
}
