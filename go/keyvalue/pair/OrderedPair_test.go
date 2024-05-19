package pair

import "testing"

func TestOrderedPair_Struct(t *testing.T) {
	t.Run("test bool:bool", func(t *testing.T) {
		var o OrderedPair[bool, bool]
		o.lock.Lock()
		o.lock.Unlock()
		if o.data != nil {
			t.Fatalf("o.data should be nil at this time")
		}
	})
	t.Run("test to ensure Key and Value can be different types (string:int)", func(t *testing.T) {
		var o OrderedPair[string, int]
		if o.data != nil {
			t.Fatalf("o.data should be nil at this time")
		}
	})
}
