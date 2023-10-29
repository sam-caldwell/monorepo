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
	//
	// Create an uninitialized state of ByteCounter
	//
	var b ByteCounter
	if b.v != nil {
		t.Fatalf("expected nil byte array initially")
	}
	//
	// Sad: Create ByteCounter with bad value
	//
	if b, err = NewByteCounter(0); err == nil {
		if err.Error() != "ByteCounter size must be > 0" {
			t.Fatalf("Expected error.  Got %s", err.Error())
		}
	}
	//
	// Sad: Create ByteCounter with bad value
	//
	if b, err = NewByteCounter(-1); err == nil {
		if err.Error() != "ByteCounter size must be > 0" {
			t.Fatalf("Expected error.  Got %s", err.Error())
		}
	}
	//
	// Happy: Create an initialized ByteCounter (1)
	//
	if b, err = NewByteCounter(1); err != nil {
		t.Fatalf("Expected nil err")
	}
	if b.v == nil {
		t.Fatal("expected non-nil byte array")
	}
	if len(b.v) != 1 {
		t.Fatal("Expected 1-element byte array")
	}
	//
	// Happy: Create an initialized ByteCounter (10)
	//
	if b, err = NewByteCounter(10); err != nil {
		t.Fatalf("Expected nil err")
	}
	if len(b.v) != 10 {
		t.Fatal("Expected 10-element byte array")
	}
}
