package simple

import (
	"github.com/sam-caldwell/monorepo/go/misc"
	"testing"
)

func TestSet_Has(t *testing.T) {

	t.Run("test int", func(t *testing.T) {
		var set Set[int]
		t.Run("uninitialized test", func(t *testing.T) {
			if set.Has(-42) {
				t.Fatal("Expected -42 to not exist in the set, but it does")
			}
		})
		t.Run("add data", func(t *testing.T) {
			set.Init()
			set.data[-1] = misc.NullObjectStruct{}
			set.data[-2] = misc.NullObjectStruct{}
			set.data[-3] = misc.NullObjectStruct{}
			set.data[-4] = misc.NullObjectStruct{}
		})
		t.Run("test the data", func(t *testing.T) {
			for _, i := range []int{-1, -2, -3, -4} {
				if !set.Has(i) {
					t.Fatal("Expected 1 to exist in the set, but it didn't")
				}
			}
		})
	})
	t.Run("test uint", func(t *testing.T) {
		var set Set[uint]
		t.Run("uninitialized test", func(t *testing.T) {
			if set.Has(42) {
				t.Fatal("Expected 42 to not exist in the set, but it does")
			}
		})
		t.Run("add data", func(t *testing.T) {
			set.Init()
			set.data[1] = misc.NullObjectStruct{}
			set.data[2] = misc.NullObjectStruct{}
			set.data[3] = misc.NullObjectStruct{}
			set.data[4] = misc.NullObjectStruct{}
		})
		t.Run("test the data", func(t *testing.T) {
			for _, i := range []int{1, 2, 3, 4} {
				if !set.Has(uint(i)) {
					t.Fatal("Expected 1 to exist in the set, but it didn't")
				}
			}
		})
	})
	t.Run("test bool", func(t *testing.T) {
		var set Set[bool]
		t.Run("uninitialized test", func(t *testing.T) {
			if set.Has(true) {
				t.Fatal("Expected true to not exist in the set, but it does")
			}
		})
		t.Run("add data", func(t *testing.T) {
			set.Init()
			set.data[true] = misc.NullObjectStruct{}
			set.data[true] = misc.NullObjectStruct{}
		})
		t.Run("test the data", func(t *testing.T) {
			if !set.Has(true) {
				t.Fatal("Expected true to exist in the set, but it does")
			}
			if set.Has(false) {
				t.Fatal("Expected false to not exist in the set, but it does")
			}
		})
	})
	t.Run("test string", func(t *testing.T) {
		var set Set[string]
		t.Run("uninitialized test", func(t *testing.T) {
			// Check if the set has no items initially
			if set.Has("apple") {
				t.Fatal("Expected 'apple' to not exist in the set, but it does")
			}
		})
		t.Run("add data", func(t *testing.T) {
			set.Init()
			set.data["apple"] = misc.NullObjectStruct{}
			set.data["orange"] = misc.NullObjectStruct{}
			set.data["banana"] = misc.NullObjectStruct{}
		})
		t.Run("test the data", func(t *testing.T) {
			if !set.Has("apple") {
				t.Fatal("Expected 'apple' to exist in the set, but it didn't")
			}
			if !set.Has("orange") {
				t.Fatal("Expected 'orange' to not exist in the set, but it does")
			}
			if !set.Has("banana") {
				t.Fatal("Expected 'banana' to not exist in the set, but it does")
			}
			if set.Has("kiwi") {
				t.Fatal("Expected 'kiwi' to not exist in the set, but it did")
			}
		})
	})
}
