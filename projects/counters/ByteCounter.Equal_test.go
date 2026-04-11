package counters

import "testing"

func TestByteCounter_Equal(t *testing.T) {
	var lhs ByteCounter
	var rhs ByteCounter

	t.Run("Happy Path: Two equal counters", func(t *testing.T) {
		_ = lhs.Set(0, 0)
		_ = rhs.Set(0, 0)
		if !lhs.Equal(&rhs) {
			t.Fatalf("failed\n"+
				"lhs: %v\n"+
				"rhs: %v",
				lhs.String(),
				rhs.String())
		}
	})

	t.Run("Happy Path: Two equal counters", func(t *testing.T) {
		_ = lhs.Set(0, 1)
		_ = rhs.Set(0, 1)
		if !lhs.Equal(&rhs) {
			t.Fatalf("failed\n"+
				"lhs: %v\n"+
				"rhs: %v",
				lhs.String(),
				rhs.String())
		}
	})

	t.Run("Happy Path: Two unequal counters pos 0", func(t *testing.T) {
		_ = lhs.Set(0, 0)
		_ = rhs.Set(0, 1)
		if lhs.Equal(&rhs) {
			t.Fatalf("failed\n"+
				"lhs: %v\n"+
				"rhs: %v",
				lhs.String(),
				rhs.String())
		}
	})

	t.Run("Happy Path: Two equal counters pos 1", func(t *testing.T) {
		_ = lhs.Set(0, 0)
		_ = rhs.Set(0, 0)
		_ = lhs.Set(1, 1)
		_ = rhs.Set(1, 1)
		if !lhs.Equal(&rhs) {
			t.Fatalf("failed\n"+
				"lhs: %v\n"+
				"rhs: %v",
				lhs.String(),
				rhs.String())
		}
	})
}
