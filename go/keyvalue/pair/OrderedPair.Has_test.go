package pair

import (
	"fmt"
	"testing"
)

func TestOrderedPair_Has(t *testing.T) {
	var o OrderedPair[string, int]
	t.Run("add data", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("test%d", i)
			value := i
			o.Add(key, value)
		}
	})
	t.Run("verify data and ordering", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("test%d", i)
			if !o.Has(key) {
				t.Fatalf("key %s should exist", key)
			}
		}
	})
}
