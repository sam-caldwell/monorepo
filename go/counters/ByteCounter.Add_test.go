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
	b, err := NewByteCounter(10)
	if err != nil {
		t.Fatal("expect no error when creating counter")
	}
	for i := 0; i < len(b.v); i++ {
		if b.v[i] != 0 {
			t.Fatal("expect zeroed array")
		}
	}
	if err := b.Add(1); err != nil {
		t.Fatalf("expect no error on add(1):%v", err)
	}
	if b.v[0] != 1 {
		t.Fatalf("zero element should be 1")
	}
	for i := 1; i < len(b.v); i++ {
		if b.v[i] != 0 {
			t.Fatal("expect zeroed array")
		}
	}
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
	b.v[0] = 0
	if err := b.Add(256); err != nil {
		t.Fatalf("expect no error on add(2):%v", err)
	}
	if b.v[0] != 0 {
		t.Fatalf("element 0 should be 0")
	}
	if b.v[1] != 1 {
		t.Fatalf("element 1 should be 1")
	}
	for i := 2; i < len(b.v); i++ {
		if b.v[i] != 0 {
			t.Fatal("expect zeroed array")
		}
	}
}
