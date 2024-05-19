package keyvalue

import (
	"github.com/google/uuid"
	"testing"
)

func TestKeyValue_ValueWidth(t *testing.T) {
	t.Run("test with nil data", func(t *testing.T) {
		var kv KeyValue[string, string]
		if width := kv.ValueWidth(); width != 0 {
			t.Fatalf("expect 0 width, got %d", width)
		}
	})
	t.Run("test with data", func(t *testing.T) {
		var kv KeyValue[string, string]
		t.Run("initialize data", func(t *testing.T) {
			kv.data = make(map[string]string)
			values := []string{
				"a",
				"ab",
				"abc",
				"abcd",
				"abcde",
				"short",
				"medium_size",
				"really_long_key",
				"even_longer_key",
				"this_key_is_the_longest",
			}
			for _, value := range values {
				key := uuid.New().String()
				kv.data[key] = value
			}
		})
		t.Run("get width", func(t *testing.T) {
			expectedWidth := len("this_key_is_the_longest")
			width := kv.ValueWidth()

			if width != expectedWidth {
				t.Fatalf("expect %d width, got %d", expectedWidth, width)
			}
		})
	})
}

//
//func TestValueWidthHappyPath(t *testing.T) {
//	kv := KeyValue{
//		data: Map{
//			"key1": "short",
//			"key2": "very long value",
//		},
//	}
//
//	width := kv.ValueWidth()
//
//	if width != len("very long value") {
//		t.Fatalf("Incorrect width. Expected %d, but got %d", len("very long value"), width)
//	}
//}
//
//func TestValueWidthWithEmptyMap(t *testing.T) {
//	kv := KeyValue{}
//
//	width := kv.ValueWidth()
//
//	if width != 0 {
//		t.Fatalf("Incorrect width. Expected 0, but got %d", width)
//	}
//}
//
//func TestValueWidthWithNilMap(t *testing.T) {
//	kv := KeyValue{}
//
//	width := kv.ValueWidth()
//
//	if width != 0 {
//		t.Fatalf("Incorrect width. Expected 0, but got %d", width)
//	}
//}
