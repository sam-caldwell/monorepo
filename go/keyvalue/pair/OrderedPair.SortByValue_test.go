package pair

import (
	"github.com/google/uuid"
	"testing"
)

func TestOrderedPair_SortByValue(t *testing.T) {
	var o OrderedPair[string, int]
	t.Run("add data", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			o.Add(uuid.New().String(), i)
		}
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("verify data and ordering", func(t *testing.T) {
		for i := 10; i < 0; i-- {
			if o.data[i].Value != i {
				t.Fatalf("expected key mismatch")
			}
		}
	})
	t.Run("Verify length before sort", func(t *testing.T) {
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("Test SortByValue", func(t *testing.T) {
		o.SortByValue()
	})
	t.Run("Verify length after sort", func(t *testing.T) {
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("Verify Sorted Order", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			expectedKey := o.data[i].Key
			expectedValue := i
			if o.data[i].Value != expectedValue {
				t.Fatalf("expected key mismatch [%d](key: '%s', value: '%s')", i, expectedKey, o.data[i].Key)
			}
		}
	})
}
