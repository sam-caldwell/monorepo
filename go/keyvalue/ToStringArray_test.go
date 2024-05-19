package keyvalue

import (
	"fmt"
	"testing"
)

func TestKeyValueToStringArray(t *testing.T) {
	t.Run("test with nil KeyValue", func(t *testing.T) {
		var kv KeyValue[string, string]
		t.Run("test with pretty:true", func(t *testing.T) {
			output := kv.ToStringArray(":", "\n", false)
			if v := fmt.Sprintf("%v", output); v != "[]" {
				t.Fatalf("expect empty result.  Got %v", output)
			}
		})
		t.Run("test with pretty:false", func(t *testing.T) {
			output := kv.ToStringArray(":", "\n", true)
			if v := fmt.Sprintf("%v", output); v != "[]" {
				t.Fatalf("expect empty result.  Got %v", output)
			}
		})
	})
	t.Run("test with sorted data and pretty:false", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		}
		expected := "[key1:value1\n key2:value2\n key3:value3\n key4:value4\n]"
		output := kv.ToStringArray(":", "\n", false)
		if v := fmt.Sprintf("%v", output); v != expected {
			t.Fatalf("expected result not found.\n\tGot\n'%v'\n\tExpected\n'%v'\n", output, expected)
		}
	})
	t.Run("test with unsorted data and pretty:false", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key4": "value4",
			"key2": "value2",
			"key3": "value3",
			"key1": "value1",
		}
		expected := "[key1:value1\n key2:value2\n key3:value3\n key4:value4\n]"
		output := kv.ToStringArray(":", "\n", false)
		if v := fmt.Sprintf("%v", output); v != expected {
			t.Fatalf("expected result not found.\n\tGot\n'%v'\n\tExpected\n'%v'\n", output, expected)
		}
	})
	t.Run("test with sorted data and pretty:true", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key1":   "value1",
			"key2":   "value2",
			"key3":   "value3",
			"key400": "value4",
		}
		expected := "[key1  :value1\n key2  :value2\n key3  :value3\n key400:value4\n]"
		output := kv.ToStringArray(":", "\n", true)
		if v := fmt.Sprintf("%v", output); v != expected {
			t.Fatalf("expected result not found.\n\tGot\n'%v'\n\tExpected\n'%v'\n", output, expected)
		}
	})
}
