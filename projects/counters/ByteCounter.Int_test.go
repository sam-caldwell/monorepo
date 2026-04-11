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
	var err error
	var b *ByteCounter

	t.Run("create a new counter", func(t *testing.T) {
		b, err = NewByteCounter(10)
		if err != nil {
			t.Fatal("expected no error")
		}
	})

	t.Run("Happy: confirm initial state (0) using.Int()", func(t *testing.T) {
		b.v = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		actual := big.Int{}
		actual.SetBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		expected := b.Int()
		t.Log("inside test 2 (Int done)")
		if actual.Cmp(expected) != 0 {
			t.Fatal("Expected empty (0) value")
		}
		t.Log("initial state confirmed.")
	})

	t.Run("Happy: Increment the state and confirm .Int() works", func(t *testing.T) {
		//t.Logf("state: %v", b.v)
		_ = b.Increment()
		//t.Logf("state: %v", b.v)
		if !bytes.Equal(b.v, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
			t.Fatalf("increment failed to change the state\nb.v:%v", b.v)
		}
		expected := big.NewInt(1)
		if expected.Cmp(b.Int()) != 0 {
			t.Fatalf("expected 0 indicating equality, got non-zero\n"+
				"expected: %v\n"+
				"actual:   %v",
				expected, b.Int())
		}
	})

	t.Run("Happy: Test Increment() a bit more.", func(t *testing.T) {
		stopValue := big.NewInt(100)
		b.v = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		for i := big.NewInt(1); i.Cmp(stopValue) < 0; i.Add(i, big.NewInt(1)) {
			_ = b.Increment()
			if i.Cmp(b.Int()) != 0 {
				t.Fatalf("mismatch values\n"+
					"i: %v\n"+
					"b: %v",
					i.Int64(), b.Int())
			}
			//t.Logf("test increment passes at %v", i)
		}
	})
}
