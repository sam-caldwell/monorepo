package hashScanner

import (
	"testing"
)

/*
 * hashScanner Worker struct test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner worker structure test
 */

func TestWorkerStruct(t *testing.T) {
	var w Worker
	if w.id != 0 {
		t.Fatal("worker id expects 0")
	}
	if w.offset != 0 {
		t.Fatal("worker offset expects 0")
	}
	if w.counter.String() != "" {
		t.Fatal("expect worker.counter not nil")
	}
}
