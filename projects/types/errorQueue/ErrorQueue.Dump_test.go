package errorQueue

import (
	"fmt"
	"testing"
)

func TestErrorQueue_Dump(t *testing.T) {
	var e ErrorQueue
	e.Push(fmt.Errorf("test1"))
	e.Push(fmt.Errorf("test2"))
	if len(e) != 2 {
		t.Fatalf("expected size mismatch")
	}
	if len(e.Dump()) != 2 {
		t.Fatalf("expected dump size mismatch")
	}
	actual := e.Dump()
	if len(actual) != 2 {
		t.Fatalf("dump size expected 2 elements")
	}
	if actual[0].Error() != "test1" {
		t.Fatalf("test1 mismatch")
	}
	if actual[1].Error() != "test2" {
		t.Fatalf("test2 mismatch")
	}
}
