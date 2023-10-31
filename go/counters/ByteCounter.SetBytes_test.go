package counters

import "testing"

/*
 * ByteCounter.SetBytes()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for ByteCounter.SetBytes()
 */

func (c *ByteCounter) TestByteCounter_SetBytes_Overflow(t *testing.T) {
	func() {
		b, err := NewByteCounter(1)
		if err != nil {
			t.Fatal("expected no error")
		}
		err = b.SetBytes(2, []byte{0x01, 0x02})
		if err == nil {
			t.Fatal("error expected")
		}
		if err.Error() != "overflow error" {
			t.Fatal("expected overflow error not found")
		}
	}()
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatal("expected no error")
	}
	err = b.SetBytes(2, []byte{0x01, 0x02})
	if err != nil {
		t.Fatal("expected no error")
	}
	if b.v[0] != 0 {
		t.Fatal("expect element 0 to be 0")
	}
	if b.v[1] != 0 {
		t.Fatal("expect element 1 to be 0")
	}
	if b.v[2] != 0x01 {
		t.Fatal("expect element 2 to be 0x01")
	}
	if b.v[3] != 0x02 {
		t.Fatal("expect element 3 to be 0x02")
	}
	for i := 4; i < len(b.v); i++ {
		if b.v[i] != 0 {
			t.Fatalf("expect element %d to be 0x00", i)
		}
	}
}
