package pair

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"reflect"
	"testing"
)

func TestOrderedPair_Get(t *testing.T) {
	t.Run("test with no data", func(t *testing.T) {
		var o OrderedPair[string, string]

		if value, err := o.Get(0); err == nil {
			if err.Error() != errors.IndexOutOfRange {
				t.Fatalf("expected index out of range, got %v", err.Error())
			}
			t.Fatal("expected error")
		} else {
			if !reflect.DeepEqual(value, Pair[string, string]{}) {
				t.Fatal("value should be empty.")
			}
		}
	})
	t.Run("test with data", func(t *testing.T) {
		var o OrderedPair[string, string]
		o.Add("key1", "value1")
		o.Add("key2", "value2")

		if value, err := o.Get(0); err != nil {
			t.Fatal(err)
		} else {
			if value.Key != "key1" || value.Value != "value1" {
				t.Fatal("expected key/value pair")
			}
		}
	})
}
