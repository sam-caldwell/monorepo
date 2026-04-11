package keyvalue

import (
	"fmt"
	"testing"
)

func TestInitialize(t *testing.T) {

	t.Run("initialize from declared variable", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.Initialize(0, false)
		kv.data["test"] = "data"
		if _, ok := kv.data["test"]; !ok {
			t.Fatalf("test data not found")
		}
	})

	t.Run("initialize with negative sz value", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.Initialize(-1024, false)
		if kv.data == nil {
			t.Fatalf("test data not initialized")
		}
		if len(kv.data) != 0 {
			t.Fatalf("Expect size 0")
		}
		kv.Initialize(-1024, true)
		if kv.data == nil {
			t.Fatalf("test data not initialized")
		}
		if len(kv.data) != 0 {
			t.Fatalf("Expect size 0")
		}
	})

	t.Run("initialize with overwrite (initial data should be removed)", func(t *testing.T) {
		var kv KeyValue[string, string]
		t.Run("setup the initial state", func(t *testing.T) {
			kv = KeyValue[string, string]{
				data: map[string]string{
					"key0": "value0",
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
			}
		})
		t.Run("use the .Initialize() method with allowOverwrite", func(t *testing.T) {
			if len(kv.data) != 6 {
				t.Fatalf("Expect size 6")
			}
			kv.Initialize(0, true)
			if kv.data == nil {
				t.Fatalf("Failed to initialize keyvalue data")
			}
			if len(kv.data) != 0 {
				t.Fatalf("Expect size 0")
			}
		})
		t.Run("verify that .Initialize() eliminated the prior state", func(t *testing.T) {
			for i := 0; i < len(kv.data); i++ {
				key := fmt.Sprintf("key%d", i)
				if _, ok := kv.data[key]; ok {
					t.Fatalf("%s exists (it should not)", key)
				}
			}
		})
		t.Run("add a new record and verify it exists", func(t *testing.T) {
			kv.data["test_key"] = "test_value"
			if v, ok := kv.data["test_key"]; !ok || v != "test_value" {
				t.Fatalf("Failed to update data")
			}
		})
	})
	t.Run("happy path: initialize with no overwrite (initial data should persist)", func(t *testing.T) {
		var kv KeyValue[string, string]
		t.Run("setup the initial state", func(t *testing.T) {
			kv = KeyValue[string, string]{
				data: map[string]string{
					"key0": "value0",
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
			}
		})
		t.Run("use the .Initialize() method with allowOverwrite", func(t *testing.T) {
			kv.Initialize(0, false)
			if kv.data == nil {
				t.Fatalf("Failed to initialize keyvalue data")
			}
		})
		t.Run("verify that .Initialize() eliminated the prior state", func(t *testing.T) {
			for i := 0; i < len(kv.data); i++ {
				key := fmt.Sprintf("key%d", i)
				if _, ok := kv.data[key]; !ok {
					t.Fatalf("%s does not exist (it should)", key)
				}
			}
		})
		t.Run("add a new record and verify it exists", func(t *testing.T) {
			kv.data["test_key"] = "test_value"
			if v, ok := kv.data["test_key"]; !ok || v != "test_value" {
				t.Fatalf("Failed to update data")
			}
		})
	})
}
