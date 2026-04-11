package pair

import "testing"

func TestOrderedPair_Count(t *testing.T) {
	t.Run("Test with nil OrderedPair", func(t *testing.T) {
		var p OrderedPair[string, string]
		if p.Count() != 0 {
			t.Fatalf("want 0, got %v", p.Count())
		}
	})
	t.Run("Test with non-empty OrderedPair", func(t *testing.T) {
		var p OrderedPair[string, string]
		p.Add("test", "data")
		if p.Count() != 1 {
			t.Fatalf("want 1, got %v", p.Count())
		}
		p.Add("test1", "data1")
		if p.Count() != 2 {
			t.Fatalf("want 2, got %v", p.Count())
		}
	})
}
