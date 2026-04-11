package errorQueue

import (
	"fmt"
	"testing"
)

func TestErrorQueue_Push(t *testing.T) {
	var e ErrorQueue
	e.Push(fmt.Errorf("test1"))
	if error(e[0]).Error() != "test1" {
		t.Fatalf("push failed")
	}
}
