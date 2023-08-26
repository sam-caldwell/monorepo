package keyvalue

import (
	"testing"
)

func TestGetInt_KeyExistsAndIsInt(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": 42,
			"key2": -10,
		},
	}

	value, err := kv.GetInt("key1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if value != 42 {
		t.Errorf("Expected value '42', but got '%v'", value)
	}
}

func TestGetInt_KeyExistsButIsNotInt(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
		},
	}

	value, err := kv.GetInt("key1")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != 0 {
		t.Errorf("Expected value '0', but got '%v'", value)
	}
}

func TestGetInt_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": 42,
		},
	}

	value, err := kv.GetInt("key2")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != 0 {
		t.Errorf("Expected value '0', but got '%v'", value)
	}
}
