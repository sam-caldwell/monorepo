package keyvalue

import (
	"testing"
)

func TestValueWidthHappyPath(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "short",
			"key2": "very long value",
		},
	}

	width := kv.ValueWidth()

	if width != len("very long value") {
		t.Errorf("Incorrect width. Expected %d, but got %d", len("very long value"), width)
	}
}

func TestValueWidthWithEmptyMap(t *testing.T) {
	kv := KeyValue{}

	width := kv.ValueWidth()

	if width != 0 {
		t.Errorf("Incorrect width. Expected 0, but got %d", width)
	}
}

func TestValueWidthWithNilMap(t *testing.T) {
	kv := KeyValue{}

	width := kv.ValueWidth()

	if width != 0 {
		t.Errorf("Incorrect width. Expected 0, but got %d", width)
	}
}
