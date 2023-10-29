package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter tests
 */

import (
	"bytes"
	"math/big"
	"testing"
)

func TestByteCounter_Int(t *testing.T) {
	//
	// setup: create a new counter
	//
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatal("expected no error")
	}
	//
	// Happy: confirm initial state (0) using.Int()
	//
	func() {
		b.v = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		actual := big.Int{}
		actual.SetBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		expected := b.Int()
		if actual.Cmp(expected) != 0 {
			t.Fatal("Expected empty (0) value")
		}
	}()
	//
	// Happy: Increment the state and confirm .Int() works
	//
	func() {
		t.Logf("state: %v", b.v)
		if err := b.Increment(); err != nil {
			t.Fatalf("error incrementing: %v", err)
		}
		t.Logf("state: %v", b.v)
		if !bytes.Equal(b.v, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
			t.Fatalf("increment failed to change the state\nb.v:%v", b.v)
		}
		expected := big.NewInt(1)
		if expected.Cmp(b.Int()) == 0 {
			t.Fatal("expected 1")
		}
	}()
	//
	// Happy: Test Increment() a bit more.
	//
	func() {
		stopValue := big.NewInt(100)
		b.v = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		for i := big.NewInt(1); i.Cmp(stopValue) < 0; i.Add(i, big.NewInt(1)) {
			if err := b.Increment(); err != nil {
				t.Fatalf("error incrementing: %v", err)
			}
			if i.Cmp(b.Int()) == 0 {
				t.Fatalf("mismatch at %v (%v)", b.v, i)
			}
			t.Logf("test increment passes at %v", i)
		}
	}()
}
