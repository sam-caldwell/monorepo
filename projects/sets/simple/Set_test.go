package simple

import "testing"

func TestSet_Struct(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var s Set[int]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("int8", func(t *testing.T) {
		var s Set[int8]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("int16", func(t *testing.T) {
		var s Set[int16]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("int32", func(t *testing.T) {
		var s Set[int32]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("int64", func(t *testing.T) {
		var s Set[int64]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uint", func(t *testing.T) {
		var s Set[uint]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uint8", func(t *testing.T) {
		var s Set[uint8]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uint16", func(t *testing.T) {
		var s Set[uint16]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uint32", func(t *testing.T) {
		var s Set[uint32]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uint64", func(t *testing.T) {
		var s Set[uint64]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("rune", func(t *testing.T) {
		var s Set[rune]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("string", func(t *testing.T) {
		var s Set[string]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("bool", func(t *testing.T) {
		var s Set[bool]
		if s.data != nil {
			t.Fatal("expect nil")
		}
	})
}
