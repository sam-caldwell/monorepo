package counters

import "testing"

/*
* ByteCounter
* (c) 2023 Sam Caldwell.  See License.txt

* ByteCounter tests
 */
func TestNewByteCounter(t *testing.T) {
	var err error
	var b ByteCounter
	if b.v != nil {
		t.Fatalf("expected nil byte array initially")
	}
	if b, err = NewByteCounter(0); err == nil {
		if err.Error() != "ByteCounter size must be > 0" {
			t.Fatalf("Expected error.  Got %s", err.Error())
		}
	}
	if b, err = NewByteCounter(-1); err == nil {
		if err.Error() != "ByteCounter size must be > 0" {
			t.Fatalf("Expected error.  Got %s", err.Error())
		}
	}
	if b, err = NewByteCounter(1); err != nil {
		t.Fatalf("Expected nil err")
	}
	if b.v == nil {
		t.Fatal("expected non-nil byte array")
	}
	if len(b.v) != 1 {
		t.Fatal("Expected 1-element byte array")
	}

	if b, err = NewByteCounter(10); err != nil {
		t.Fatalf("Expected nil err")
	}
	if len(b.v) != 10 {
		t.Fatal("Expected 10-element byte array")
	}
}

func TestByteCounter_Increment(t *testing.T) {
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
	for i := 0; i < len(b.v); i++ {
		b.v[i] = 255
	}
	if err := b.Increment(); err != nil {
		if err.Error() != "overflow error" {
			t.Fatal("Expected overflow error")
		}
	}
}
