package simple

import (
	"testing"
)

func TestSet_Delete(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		var set Set[int]
		t.Run("Add data", func(t *testing.T) {
			_ = set.Add(1)
			_ = set.Add(2)
			_ = set.Add(3)
			_ = set.Add(4)
			_ = set.Add(5)
			_ = set.Add(6)
			_ = set.Add(7)
		})
		t.Run("verify data[3] exists and has() method works", func(t *testing.T) {
			if !set.Has(3) {
				t.Fatal("record data[3] does not exist or Has() method fails")
			}
		})
		t.Run("Delete data[3]", func(t *testing.T) {
			set.Delete(3)
		})
		t.Run("verify data[3] exists and has() method works", func(t *testing.T) {
			if set.Has(3) {
				t.Fatal("record data[3] should have been deleted, but it still exists")
			}
		})
	})
	t.Run("test int", func(t *testing.T) {
		var set Set[string]
		t.Run("Add data", func(t *testing.T) {
			_ = set.Add("1")
			_ = set.Add("2")
			_ = set.Add("3")
			_ = set.Add("4")
			_ = set.Add("5")
			_ = set.Add("6")
			_ = set.Add("7")
		})
		t.Run("verify data[3] exists and has() method works", func(t *testing.T) {
			if !set.Has("3") {
				t.Fatal("record data[3] does not exist or Has() method fails")
			}
		})
		t.Run("Delete data[3]", func(t *testing.T) {
			set.Delete("3")
		})
		t.Run("verify data[3] exists and has() method works", func(t *testing.T) {
			if set.Has("3") {
				t.Fatal("record data[3] should have been deleted, but it still exists")
			}
		})
		t.Run("delete non-existent record", func(t *testing.T) {
			set.Delete("this doesn't exist")
		})
	})
}
