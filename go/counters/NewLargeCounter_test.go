package counters

import "testing"

func TestNewLargeCounter(t *testing.T) {
	func() {
		var o LargeCounter
		if o != nil {
			t.Fatal("LargeCounter expects nil initially")
		}
	}()
	func() {
		o, err := NewLargeCounter(64)
		if err != nil {
			t.Fatalf("expect no error. err:%v", err)
		}
		if len(*o) != 1 {
			t.Fatalf("expect 1-element counter: %d", len(*o))
		}
	}()
	func() {
		o, err := NewLargeCounter(65)
		if err != nil {
			t.Fatalf("expect no error. err:%v", err)
		}
		if len(*o) != 2 {
			t.Fatalf("expect 2-element counter: %d", len(*o))
		}
	}()
	func() {
		o, err := NewLargeCounter(66)
		if err != nil {
			t.Fatalf("expect no error. err:%v", err)
		}
		if len(*o) != 2 {
			t.Fatalf("expect 2-element counter: %d", len(*o))
		}
	}()
	func() {
		o, err := NewLargeCounter(128)
		if err != nil {
			t.Fatalf("expect no error. err:%v", err)
		}
		if len(*o) != 2 {
			t.Fatalf("expect 2-element counter: %d", len(*o))
		}
	}()
	func() {
		o, err := NewLargeCounter(129)
		if err != nil {
			t.Fatalf("expect no error. err:%v", err)
		}
		if len(*o) != 3 {
			t.Fatalf("expect 2-element counter: %d", len(*o))
		}
	}()
	func() {
		o, err := NewLargeCounter(0)
		if o != nil {
			t.Fatal("expect nil counter")
		}
		if err == nil {
			t.Fatal("expected an error")
		}
		if err.Error() != "bounds check error" {
			t.Fatal("expect bounds check error")
		}

	}()
	func() {
		o, err := NewLargeCounter(-1)
		if o != nil {
			t.Fatal("expect nil counter")
		}
		if err == nil {
			t.Fatal("expected an error")
		}
		if err.Error() != "bounds check error" {
			t.Fatal("expect bounds check error")
		}

	}()
}
