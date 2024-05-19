package pair

import (
	"fmt"
	"testing"
)

func TestOrderedPair_Walk(t *testing.T) {
	walkerFunc := func(actualKey string, actualValue int) error {
		expectedKey := fmt.Sprintf("test%d", actualValue)
		if actualKey != expectedKey {
			return fmt.Errorf("key mismatch: %s", actualKey)
		}
		return nil
	}
	var o OrderedPair[string, int]
	t.Run("add data", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			o.Add(fmt.Sprintf("test%d", i), i)
		}
		if len(o.data) != 10 {
			t.Fatalf("expected length 10 got %d", len(o.data))
		}
	})
	t.Run("walk the data", func(t *testing.T) {
		if err := o.Walk(walkerFunc); err != nil {
			t.Fatal(err)
		}
	})
}
