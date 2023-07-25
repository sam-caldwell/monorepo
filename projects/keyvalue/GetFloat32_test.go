package keyvalue

import (
	"testing"
)

func TestGetFloat32_KeyExistsAndIsFloat32(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": float32(3.14),
			"key2": float32(2.71828),
		},
	}

	value, err := kv.GetFloat32("key1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if value != float32(3.14) {
		t.Errorf("Expected value '3.14', but got '%v'", value)
	}
}

func TestGetFloat32_KeyExistsButIsNotFloat32(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
		},
	}

	value, err := kv.GetFloat32("key1")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != float32(0.0) {
		t.Errorf("Expected value '0.0', but got '%v'", value)
	}
}

func TestGetFloat32_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": float32(3.14),
		},
	}

	value, err := kv.GetFloat32("key2")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != float32(0.0) {
		t.Errorf("Expected value '0.0', but got '%v'", value)
	}
}
