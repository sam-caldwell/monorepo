package keyvalue

import (
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("test with string:bool pair", func(t *testing.T) {
		kv := KeyValue[string, bool]{
			data: map[string]bool{
				"key1": true,
				"key2": false,
			},
		}
		t.Run("test with value: true", func(t *testing.T) {
			if value, err := kv.Get("key1"); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if !value {
					t.Fatalf("Expected value to be true, but got %v", value)
				}
			}
		})
		t.Run("test with value: false", func(t *testing.T) {
			if value, err := kv.Get("key2"); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if value {
					t.Fatalf("Expected value to be false, but got %v", value)
				}
			}
		})
	})
	t.Run("test with int:bool pair", func(t *testing.T) {
		kv := KeyValue[int, bool]{
			data: map[int]bool{
				1: true,
				2: false,
			},
		}
		t.Run("test with value: true", func(t *testing.T) {
			if value, err := kv.Get(1); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if !value {
					t.Fatalf("Expected value to be true, but got %v", value)
				}
			}
		})
		t.Run("test with value: false", func(t *testing.T) {
			if value, err := kv.Get(2); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if value {
					t.Fatalf("Expected value to be false, but got %v", value)
				}
			}
		})
	})
	t.Run("test with int:float32 pair", func(t *testing.T) {
		kv := KeyValue[int, float32]{
			data: map[int]float32{
				1: 1.12,
				2: 3.1415,
			},
		}
		t.Run("test with value1: 1.12", func(t *testing.T) {
			if value, err := kv.Get(1); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if value != 1.12 {
					t.Fatalf("Expected value to be 1.12, but got %v", value)
				}
			}
		})
		t.Run("test with value2: 3.1415", func(t *testing.T) {
			if value, err := kv.Get(2); err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			} else {
				if value != 3.1415 {
					t.Fatalf("Expected value to be 3.1415, but got %v", value)
				}
			}
		})
	})
}
