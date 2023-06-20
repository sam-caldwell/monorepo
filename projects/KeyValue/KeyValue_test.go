package keyvalue

import (
	"testing"
)

func TestKeyValueInitialization(t *testing.T) {
	kv := KeyValue{}

	if kv.data != nil {
		t.Errorf("Expected data to be nil, but got %+v", kv.data)
	}
}

func TestKeyValueDataInitialization(t *testing.T) {
	kv := KeyValue{
		data: Map{"key": "value"},
	}

	if kv.data == nil {
		t.Errorf("Expected data not to be nil")
	} else if len(kv.data) != 1 {
		t.Errorf("Expected data to have one key-value pair, but got %+v", kv.data)
	} else if kv.data["key"] != "value" {
		t.Errorf("Expected key to have value 'value', but got %+v", kv.data["key"])
	}
}

func TestKeyValueDataWithNilValue(t *testing.T) {
	kv := KeyValue{
		data: Map{"key": nil},
	}

	if kv.data == nil {
		t.Errorf("Expected data not to be nil")
	} else if len(kv.data) != 1 {
		t.Errorf("Expected data to have one key-value pair, but got %+v", kv.data)
	} else if kv.data["key"] != nil {
		t.Errorf("Expected key to have value nil, but got %+v", kv.data["key"])
	}
}
