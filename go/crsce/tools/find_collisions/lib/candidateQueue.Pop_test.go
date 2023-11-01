package hashScanner

import "testing"

/*
 * hashScanner CandidateQueue.Pop() tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate queue Pop method tests.
 */

func TestCandidateQueue_Pop(t *testing.T) {
	var o CandidateQueue
	o.Push("1", "1")
	if o.Count() != 1 {
		t.Fatal("count expected 1 after Push")
	}
	o.Push("2", "2")
	n := o.Pop()
	if o.Count() != 1 {
		t.Fatal("count expected 1 after Pop")
	}
	if n.raw != "1" {
		t.Fatal("expected record 1")
	}
	if n.hash != "1" {
		t.Fatal("expected record 1")
	}
	n = o.Pop()
	if n.raw != "2" {
		t.Fatal("expect record 2")
	}
	if n.hash != "2" {
		t.Fatal("expect record 2")
	}
	if o.Count() != 0 {
		t.Fatal("expect 0 records remain")
	}
}
