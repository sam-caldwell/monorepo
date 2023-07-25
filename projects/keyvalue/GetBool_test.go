package keyvalue

import (
	"testing"
)

func TestGetBool_KeyExistsAndIsBool(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": true,
			"key2": false,
		},
	}

	value, err := kv.GetBool("key1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if value != true {
		t.Errorf("Expected value 'true', but got '%v'", value)
	}
}

func TestGetBool_KeyExistsButIsNotBool(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
		},
	}

	value, err := kv.GetBool("key1")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != false {
		t.Errorf("Expected value 'false', but got '%v'", value)
	}
}

func TestGetBool_KeyDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": true,
		},
	}

	value, err := kv.GetBool("key2")

	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	if value != false {
		t.Errorf("Expected value 'false', but got '%v'", value)
	}
}
