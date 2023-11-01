package hashScanner

import "testing"

/*
 * hashScanner CandidateQueue Push() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure Push() method.
 */

func TestCandidateQueue_Push(t *testing.T) {
	var o CandidateQueue
	o.Push("1", "1")
	if o.Count() != 1 {
		t.Fatal("count expected 1 after Push")
	}
	o.Push("2", "2")
	if o.Count() != 2 {
		t.Fatal("count expected 2 after Push")
	}
}
