package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestKeyValue_Walk(t *testing.T) {
	t.Run("Walk with nil data", func(t *testing.T) {
		var kv KeyValue[string, string]
		fn := func(key string, value string) error {
			return nil
		}
		if kv.data != nil {
			t.Fatal("expected nil data")
		}
		err := kv.Walk(fn)
		if err == nil {
			t.Fatalf("walk should return an error")
		}
		if err.Error() != errors.UninitializedValue {
			t.Fatalf("expected UninitializedValue")
		}
	})
	t.Run("Walk with data", func(t *testing.T) {
		var kv KeyValue[string, string]
		kv.data = map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
			"key5": "value5",
		}
		count := 0
		fn := func(key string, value string) error {
			count++
			return nil
		}
		err := kv.Walk(fn)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if count != 5 {
			t.Fatalf("expected 5 items, got %d", count)
		}
	})
}
