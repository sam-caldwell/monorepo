package pair

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestOrderedPair_Add(t *testing.T) {
	t.Run("OrderedPair test with ascending sequential data", func(t *testing.T) {
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
			for i := 0; i < 10; i++ {
				expectedKey := fmt.Sprintf("test%d", i)
				expectedValue := i
				if o.data[i].Key != expectedKey {
					t.Fatalf("expected key mismatch")
				}
				if o.data[i].Value != expectedValue {
					t.Fatalf("expected value mismatch")
				}
			}
		})
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("OrderedPair test with descending sequential data", func(t *testing.T) {
		var o OrderedPair[string, int]
		t.Run("add data", func(t *testing.T) {
			for i := 10; i < 0; i-- {
				key := fmt.Sprintf("test%d", i)
				value := i
				o.Add(key, value)
			}
		})
		t.Run("verify data and ordering", func(t *testing.T) {
			for i := 10; i < 0; i-- {
				expectedKey := fmt.Sprintf("test%d", i)
				expectedValue := i
				if o.data[i].Key != expectedKey {
					t.Fatalf("expected key mismatch")
				}
				if o.data[i].Value != expectedValue {
					t.Fatalf("expected value mismatch")
				}
			}
		})
	})
	t.Run("OrderedPair test with random data (uuid)", func(t *testing.T) {
		var o OrderedPair[uuid.UUID, uuid.UUID]
		var data []uuid.UUID
		t.Run("add data", func(t *testing.T) {
			for i := 10; i < 0; i-- {
				u := uuid.New()
				o.Add(u, u)
				data = append(data, u)
			}
		})
		t.Run("verify data and ordering", func(t *testing.T) {
			for i := 10; i < 0; i-- {
				if o.data[i].Key != data[i] {
					t.Fatalf("expected key mismatch")
				}
				if o.data[i].Value != data[i] {
					t.Fatalf("expected value mismatch")
				}
			}
		})
	})

}
