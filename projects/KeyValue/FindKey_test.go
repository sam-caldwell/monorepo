package keyvalue

import (
	"testing"
)

func TestFindKey_KeyExists(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	value, found := kv.FindKey("key1")

	if !found {
		t.Errorf("Expected key 'key1' to be found, but it was not found")
	}

	if value != "value1" {
		t.Errorf("Expected value 'value1', but got '%v'", value)
	}
}

func TestFindKey_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	value, found := kv.FindKey("key3")

	if found {
		t.Errorf("Expected key 'key3' to not be found, but it was found")
	}

	if value != nil {
		t.Errorf("Expected value to be nil, but got '%v'", value)
	}
}
