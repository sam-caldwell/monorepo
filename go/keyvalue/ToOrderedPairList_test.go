package keyvalue

import "testing"

func TestKeyValue_ToOrderedPairList(t *testing.T) {
	t.Run("test with nil keyvalue", func(t *testing.T) {
		var kv KeyValue[string, string]
		if kv.data != nil {
			t.Fatalf("expecting nil data got %v", kv.data)
		}
		p := kv.ToOrderedPairList(false)
		if p.Count() != 0 {
			t.Fatalf("expecting 0 got %v", p.Count())
		}
	})
	t.Run("test with initialized keyvalue (unsorted)", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		p := kv.ToOrderedPairList(false)
		if p.Count() != 3 {
			t.Fatalf("expecting 3 got %v", p.Count())
		}
		if !p.Has("key1") {
			t.Fatalf("expecting key1")
		}
		if !p.Has("key2") {
			t.Fatalf("expecting key2")
		}
		if !p.Has("key3") {
			t.Fatalf("expecting key3")
		}
	})
	t.Run("test with initialized keyvalue (sorted)", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key2": "value2",
			"key1": "value1",
			"key4": "value4",
			"key3": "value3",
		}
		p := kv.ToOrderedPairList(true)
		if p.Count() != 4 {
			t.Fatalf("expecting 4 got %v", p.Count())
		}
		value, err := p.Get(0)
		if err != nil {
			t.Fatal(err)
		}
		if value.Key != "key1" {
			t.Fatalf("expecting Key1 got %v", value.Key)
		}
		if value.Value != "value1" {
			t.Fatalf("expecting value1 got %v", value.Value)
		}
	})
}
