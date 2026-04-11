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

func TestNewByteCounter(t *testing.T) {
	var err error
	var b *ByteCounter
	t.Run("Ensure the initial state is nil", func(t *testing.T) {
		if b != nil {
			t.Fatalf("expected nil byte array initially")
		}
	})
	t.Run("Sad: Create ByteCounter with bad value", func(t *testing.T) {
		if b, err = NewByteCounter(0); err == nil {
			if err.Error() != "ByteCounter size must be > 0" {
				t.Fatalf("Expected error.  Got %s", err.Error())
			}
		}
	})
	t.Run("Sad: Create ByteCounter with bad value", func(t *testing.T) {
		if b, err = NewByteCounter(-1); err == nil {
			if err.Error() != "ByteCounter size must be > 0" {
				t.Fatalf("Expected error.  Got %s", err.Error())
			}
		}
	})
	t.Run("Happy: Create an initialized ByteCounter (1)", func(t *testing.T) {
		for i := 1; i < 10; i++ {
			if b, err = NewByteCounter(i); err != nil {
				t.Fatalf("Expected nil err.  Got %v", err)
			}
			if b.v == nil {
				t.Fatal("expected non-nil byte array")
			}
			if len(b.v) != i {
				t.Fatalf("expected %d-element byte array", i)
			}
		}
	})
}
