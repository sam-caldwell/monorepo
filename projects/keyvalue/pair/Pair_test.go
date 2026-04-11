package pair

import "testing"

func TestOrderedPair_Pair(t *testing.T) {
	t.Run("test ordered pair 1 (basic structure)", func(t *testing.T) {
		var p Pair[int, int]
		if p.Key != 0 {
			t.Fatal("expect initial state 0")
		}
		if p.Value != 0 {
			t.Fatal("expect initial state 0")
		}
	})
	t.Run("test ordered pair 2 (different key-value types)", func(t *testing.T) {
		var p Pair[string, int]
		if p.Key != "" {
			t.Fatal("expect initial state ''")
		}
		if p.Value != 0 {
			t.Fatal("expect initial state 0")
		}
	})
}
