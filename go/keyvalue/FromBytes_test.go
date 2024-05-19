package keyvalue

import (
	"reflect"
	"testing"
)

func TestFromBytes(t *testing.T) {
	kv := KeyValue[string, string]{
		data: map[string]string{},
	}
	expectedData := KeyValue[string, string]{
		data: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	data := []byte("key1:value1\nkey2:value2\nkey3:value3")
	lineEnding := "\n"
	columnDelimiter := ":"

	if err := kv.FromBytes(&data, lineEnding, columnDelimiter); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(kv.data, expectedData.data) {
		t.Fatalf("Parsed data does not match expected data")
	}
}
