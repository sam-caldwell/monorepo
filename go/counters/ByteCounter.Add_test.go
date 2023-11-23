package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter tests
 */

import (
	"testing"
)

func TestByteCounter_Add(t *testing.T) {
	var err error
	var b *ByteCounter
	t.Run("test 1: start from zero, add one then two", func(t *testing.T) {
		t.Run("Create a new counter", func(t *testing.T) {
			if b, err = NewByteCounter(10); err != nil {
				t.Fatal("expect no error when creating counter")
			}
		})
		t.Run("clear the array state", func(t *testing.T) {
			for i := 0; i < len(b.v); i++ {
				if b.v[i] != 0 {
					t.Fatal("expect zeroed array")
				}
			}
		})
		t.Run("Add 1 to counter", func(t *testing.T) {
			if err := b.Add(1); err != nil {
				t.Fatalf("expect no error on add(1):%v", err)
			}
			if b.v[0] != 1 {
				t.Fatalf("element zero(0) should be 1")
			}
			for i := 1; i < len(b.v); i++ {
				if b.v[i] != 0 {
					t.Fatal("expect zeroed array for elements 1...n")
				}
			}
		})
		t.Run("Add 2 to counter", func(t *testing.T) {
			if err := b.Add(2); err != nil {
				t.Fatalf("expect no error on add(2):%v", err)
			}
			if b.v[0] != 3 {
				t.Fatalf("zero element should be 3. got: %v", b.v[0])
			}
			for i := 1; i < len(b.v); i++ {
				if b.v[i] != 0 {
					t.Fatal("expect zeroed array")
				}
			}
		})
	})
	t.Run("test 2: start from byte(0) 255 and add one expecting carry", func(t *testing.T) {
		t.Run("Create a new counter", func(t *testing.T) {
			if b, err = NewByteCounter(10); err != nil {
				t.Fatal("expect no error when creating counter")
			}
		})
		t.Run("clear the array state", func(t *testing.T) {
			for i := 0; i < len(b.v); i++ {
				if b.v[i] != 0 {
					t.Fatal("expect zeroed array")
				}
			}
		})
		t.Run("Set byte(0) to 255 then add 1 to observe carry", func(t *testing.T) {
			b.v[0] = 255
			if len(b.v) != 10 {
				t.Fatalf("expect length 10")
			}
			if b.v[0] != 255 {
				t.Fatalf("element 0 should be 255")
			}
			if err := b.Add(1); err != nil {
				t.Fatalf("expect no error on add(2):%v", err)
			}
			if b.v[0] != 0 {
				t.Fatalf("element 0 should be 0 got %d", b.v[0])
			}
			if b.v[1] != 1 {
				t.Fatalf("element 1 should be 1 got %d", b.v[0])
			}
			for i := 2; i < len(b.v); i++ {
				if b.v[i] != 0 {
					t.Fatal("expect zeroed array")
				}
			}
		})

	})

}
