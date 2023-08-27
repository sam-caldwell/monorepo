package keyvalue

import (
	"testing"
)

func TestSetBoolHappyPath(t *testing.T) {
	kv := KeyValue{}

	kv.SetBool("testkey", true)

	if val, ok := kv.data["testkey"].(bool); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != true {
		t.Errorf("Incorrect value. Expected true, but got %v", val)
	}
}

func TestSetBoolOverwriteExistingKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"testkey": false,
		},
	}

	kv.SetBool("testkey", true)

	if val, ok := kv.data["testkey"].(bool); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != true {
		t.Errorf("Incorrect value. Expected true, but got %v", val)
	}
}

func TestSetBoolWithNilMap(t *testing.T) {
	kv := KeyValue{}

	kv.SetBool("testkey", true)

	if val, ok := kv.data["testkey"].(bool); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != true {
		t.Errorf("Incorrect value. Expected true, but got %v", val)
	}
}
