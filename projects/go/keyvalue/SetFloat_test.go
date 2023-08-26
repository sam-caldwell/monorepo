package keyvalue

import (
	"testing"
)

func TestSetFloatHappyPath(t *testing.T) {
	kv := KeyValue{}

	kv.SetFloat("testkey", 123.456)

	if val, ok := kv.data["testkey"].(float64); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}

func TestSetFloatOverwriteExistingKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"testkey": 654.321,
		},
	}

	kv.SetFloat("testkey", 123.456)

	if val, ok := kv.data["testkey"].(float64); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}

func TestSetFloatWithNilMap(t *testing.T) {
	kv := KeyValue{}

	kv.SetFloat("testkey", 123.456)

	if val, ok := kv.data["testkey"].(float64); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != 123.456 {
		t.Errorf("Incorrect value. Expected 123.456, but got %v", val)
	}
}
