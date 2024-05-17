package pair

import "testing"

func TestCompareKey(t *testing.T) {
	t.Run("Test with boolean values", func(t *testing.T) {
		lhs := true
		rhs := true
		if v := CompareKey[bool](lhs, rhs); v != 0 {
			t.Errorf("CompareKey(lhs,rhs)=%d, want %d", v, 0)
		}
	})
	t.Run("Test with numeric values", func(t *testing.T) {

	})
	t.Run("Test with string values", func(t *testing.T) {

	})
	t.Run("Test with boolean pointers", func(t *testing.T) {

	})
	t.Run("Test with numeric pointers", func(t *testing.T) {

	})
	t.Run("Test with string pointers", func(t *testing.T) {

	})
}
