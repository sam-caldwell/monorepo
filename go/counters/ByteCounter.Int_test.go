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
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatal("expected no error")
	}
	//
	// Happy: confirm initial state from .Bytes()
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
	func() {
		b.v = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		actual := big.Int{}
		expected := b.Int()
		if actual.Cmp(expected) != 0 {
			t.Fatal("Expected empty (0) value")
		}
		t.Logf("state: %v", b.v)
		if err := b.Increment(); err != nil {
			t.Fatalf("error incrementing: %v", err)
		}
		t.Logf("state: %v", b.v)
		if bytes.Equal(b.v, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
			t.Fatalf("increment failed to change the state")
		}
		actual = *actual.Add(&actual, big.NewInt(1))
		expected = b.Int()
		if bytes.Equal(b.v, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
			t.Fatalf("actual should not be zero (%v):%v", expected, b.v)
		}
		//if actual.Cmp(expected) != 0 {
		//	t.Fatalf(
		//		"expected value (%v). actual:%v",
		//		expected.String(), actual.String())
		//}
	}()
}
