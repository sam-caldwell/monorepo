package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter tests
 */

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestByteCounter_Increment(t *testing.T) {
	//
	// Happy: Create an initialized ByteCounter (10) and verify initial state
	//
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatalf("expected ByteCounter with no errors")
	}
	if len(b.v) != 10 {
		t.Fatalf("expected 10-element byte array")
	}
	for i := 0; i < 10; i++ {
		if b.v[i] != 0 {
			t.Fatalf("expect element %d to be zero", i)
		}
	}
	//
	// Happy: increment byte 0 to 255 and confirm state
	//
	for v := byte(1); v < byte(255); v++ {
		if err := b.Increment(); err != nil {
			t.Fatalf("expect increment without error")
		}
		if b.v[0] != v {
			t.Fatal("expected zeroth element incremented")
		}
		if b.v[0] != v {
			t.Fatalf("expect element 0 to be %d", v)
		}
		for i := 1; i < 10; i++ {
			if b.v[i] != 0 {
				t.Fatalf("expect element %d to be 0", i)
			}
		}
	}
	//
	// Happy: Set byte 0 to 255 and test carry to byte 1 if incremented
	//
	b.v[0] = 255
	if err := b.Increment(); err != nil {
		t.Fatal("expect no error")
	}
	if b.v[0] != 0 {
		t.Fatalf("expect element 0 to rollover to 0")
	}
	if b.v[1] != 1 {
		t.Fatalf("expect element 1 to carry over to 1")
	}
	//
	// Happy: Set bytes 0, 1 to 255 and test carry if byte 0 is incremented.
	//
	b.v[0] = 255
	b.v[1] = 255
	if err := b.Increment(); err != nil {
		t.Fatal("expect no error")
	}
	if b.v[0] != 0 {
		t.Fatalf("expect element 0 to rollover to 0")
	}
	if b.v[1] != 0 {
		t.Fatalf("expect element 1 to carry and rollover to 0")
	}
	if b.v[2] != 1 {
		t.Fatalf("expect element 2 to carry over to 1")
	}
	//
	// Sad: Set all bytes to 255 and increment, expecting an over-flow error
	//
	for i := 0; i < len(b.v); i++ {
		b.v[i] = 255
	}
	if err := b.Increment(); err != nil {
		if err.Error() != errors.OverflowError {
			t.Fatal("Expected overflow error")
		}
	}
}
