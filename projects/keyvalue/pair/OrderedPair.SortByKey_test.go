package pair

import (
	"fmt"
	"testing"
)

func TestOrderedPair_SortByKey(t *testing.T) {
	var o OrderedPair[string, int]
	t.Run("add data", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			o.Add(fmt.Sprintf("test%d", i), i)
		}
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("verify data and ordering", func(t *testing.T) {
		for i := 10; i < 0; i-- {
			expectedKey := fmt.Sprintf("test%d", i)
			if o.data[i].Key != expectedKey {
				t.Fatalf("expected key mismatch")
			}
		}
	})
	t.Run("Verify length before sort", func(t *testing.T) {
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("Test SortByKey", func(t *testing.T) {
		o.SortByKey()
	})
	t.Run("Verify length after sort", func(t *testing.T) {
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("Verify Sorted Order", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			expectedKey := fmt.Sprintf("test%d", i)
			expectedValue := i
			if o.data[i].Key != expectedKey {
				t.Fatalf("expected key mismatch [%d](key: '%s', value: '%s')", i, expectedKey, o.data[i].Key)
			}
			if o.data[i].Value != expectedValue {
				t.Fatalf("expected key mismatch [%d](key: '%s', value: '%s')", i, expectedKey, o.data[i].Key)
			}
		}
	})
}
