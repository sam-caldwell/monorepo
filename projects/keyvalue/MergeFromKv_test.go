package keyvalue

import (
	"testing"
)

func TestMergeFromKv(t *testing.T) {
	// Initialize two keyvalue objects with some data
	var kv1 KeyValue[string, string]
	var kv2 KeyValue[string, string]
	var expected KeyValue[string, string]

	t.Run("Initialize data", func(t *testing.T) {
		kv1.data = map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		kv2.data = map[string]string{
			"key2": "value2new",
			"key3": "value3",
			"key4": "value4",
		}
		// Create a Map with the expected result after the merge
		expected.data = map[string]string{
			"key1": "value1",
			"key2": "value2new",
			"key3": "value3",
			"key4": "value4",
		}
	})
	t.Run("Call MergeFromKv()", func(t *testing.T) {
		kv1.MergeFromKv(&kv2)
	})
	t.Run("Check if the data in kv1 matches the expected result", func(t *testing.T) {
		for expectedKey, expectedValue := range expected.data {
			if actualValue, ok := kv1.data[expectedKey]; !ok {
				t.Fatalf("key %s not found in kv1", expectedKey)
			} else {
				if actualValue != expectedValue {
					t.Fatalf("key %s found in kv1 but value is wrong", expectedKey)
				}
			}
		}
	})
}
