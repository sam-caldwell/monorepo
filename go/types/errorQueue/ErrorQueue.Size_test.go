package errorQueue

import (
	"fmt"
	"testing"
)

func TestErrorQueueSize(t *testing.T) {
	var e ErrorQueue
	if e.Size() != 0 {
		t.Fatalf("expected size 0")
	}
	e.Push(fmt.Errorf("test1"))
	if e.Size() != 1 {
		t.Fatalf("expected size 1")
	}
	e.Push(fmt.Errorf("test2"))
	if e.Size() != 2 {
		t.Fatalf("expected size 2")
	}
}
