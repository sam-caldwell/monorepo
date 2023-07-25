package keyvalue

import (
	"testing"
)

func TestSetFloat32HappyPath(t *testing.T) {
	kv := KeyValue{}

	kv.SetFloat32("test_key", 123.456)

	if val, ok := kv.data["test_key"].(float32); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}

func TestSetFloat32OverwriteExistingKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"test_key": 654.321,
		},
	}

	kv.SetFloat32("test_key", 123.456)

	if val, ok := kv.data["test_key"].(float32); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}

func TestSetFloat32WithNilMap(t *testing.T) {
	kv := KeyValue{}

	kv.SetFloat32("test_key", 123.456)

	if val, ok := kv.data["test_key"].(float32); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}
