package keyvalue

import (
	"testing"
)

func TestSetIntHappyPath(t *testing.T) {
	kv := KeyValue{}

	kv.SetInt("testkey", 123)

	if val, ok := kv.data["testkey"].(int); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123 {
		t.Errorf("Incorrect value. Expected 123, but got %v", val)
	}
}

func TestSetIntOverwriteExistingKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"testkey": 654,
		},
	}

	kv.SetInt("testkey", 123)

	if val, ok := kv.data["testkey"].(int); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123 {
		t.Errorf("Incorrect value. Expected 123, but got %v", val)
	}
}

func TestSetIntWithNilMap(t *testing.T) {
	kv := KeyValue{}

	kv.SetInt("testkey", 123)

	if val, ok := kv.data["testkey"].(int); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123 {
		t.Errorf("Incorrect value. Expected 123, but got %v", val)
	}
}
