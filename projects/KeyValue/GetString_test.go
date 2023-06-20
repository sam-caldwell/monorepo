package keyvalue

import (
	"testing"
)

func TestGetString_KeyExists(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	value, err := kv.GetString("key1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if value != "value1" {
		t.Errorf("Expected value 'value1', but got '%v'", value)
	}
}

func TestGetString_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
		},
	}

	value, err := kv.GetString("key2")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != "" {
		t.Errorf("Expected value '', but got '%v'", value)
	}
}
