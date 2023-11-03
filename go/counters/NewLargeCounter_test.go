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
		o, err := NewLargeCounter(10)
		if err != nil {
			t.Fatal("expect no error")
		}
		if len(*o) != 10 {
			t.Fatal("expect 10-element counter")
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
