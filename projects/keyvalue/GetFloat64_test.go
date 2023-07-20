package keyvalue

import (
	"testing"
)

func TestGetFloat64_KeyExistsAndIsFloat64(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": 3.14,
			"key2": 2.71828,
		},
	}

	value, err := kv.GetFloat64("key1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if value != 3.14 {
		t.Errorf("Expected value '3.14', but got '%v'", value)
	}
}

func TestGetFloat64_KeyExistsButIsNotFloat64(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
		},
	}

	value, err := kv.GetFloat64("key1")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != 0.0 {
		t.Errorf("Expected value '0.0', but got '%v'", value)
	}
}

func TestGetFloat64_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": 3.14,
		},
	}

	value, err := kv.GetFloat64("key2")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != 0.0 {
		t.Errorf("Expected value '0.0', but got '%v'", value)
	}
}
