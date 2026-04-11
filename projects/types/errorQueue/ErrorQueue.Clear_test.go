package errorQueue

import (
	"fmt"
	"testing"
)

func TestErrorQueueClear(t *testing.T) {
	var e ErrorQueue
	e.Push(fmt.Errorf("test1"))
	e.Push(fmt.Errorf("test2"))

	if len(e.Dump()) != 2 {
		t.Fatalf("expected array length mismatch")
	}
	e.Clear()
	if len(e.Dump()) != 0 {
		t.Fatalf("expected array length mismatch")
	}
}
