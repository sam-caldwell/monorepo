package errorQueue

import (
	"fmt"
	"testing"
)

func TestErrorQueue(t *testing.T) {
	var e ErrorQueue
	e.Push(fmt.Errorf("test1"))
	e.Push(fmt.Errorf("test2"))

	if len(e.Dump()) != 2 {
		t.Fatalf("expected array length mismatch")
	}
	if e.Size() != len(e.Dump()) {
		t.Fatal("size mismatch")
	}
	e.Clear()
	if e.Size() != 0 {
		t.Fatal("size mismatch")
	}
}
