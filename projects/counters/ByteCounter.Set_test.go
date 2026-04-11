package counters

import "testing"

/*
 * ByteCounter.Set
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test for the ByteCounter.Set() method.
 */

func TestByteCounter_Set_NoInit(t *testing.T) {
	var b ByteCounter
	if b.v != nil {
		t.Fatal("expect nil array initially")
	}
	if err := b.Set(0, 1); err != nil {
		t.Fatalf("expect no error(1): %v", err)
	}
	if b.v == nil {
		t.Fatal("expected initialization by Set()")
	}
	if len(b.v) != 1 {
		t.Fatal("expected counter to initialize")
	}
	if b.v[0] != 1 {
		t.Fatal("expected value not set")
	}
	if err := b.Set(1, 1); err == nil {
		t.Fatal("Expected error")
	} else if err.Error() != "index error" {
		t.Fatal("expected error not found.")
	}
}

func TestByteCounter_Set_WithInit(t *testing.T) {
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatal("expected no error")
	}
	if b.v == nil {
		t.Fatal("expect b.v initialized")
	}
	if err := b.Set(2, 2); err != nil {
		t.Fatal("expected no error")
	}
	if b.v[2] != 2 {
		t.Fatal("expected value mismatch")
	}
}
