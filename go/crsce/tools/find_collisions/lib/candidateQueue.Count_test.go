package hashScanner

import "testing"

/*
 * hashScanner candidateQueue.Count() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidateQueue Count method tests.
 */

func TestCandidateQueue_Count(t *testing.T) {
	var o CandidateQueue
	o.queue = make(chan Candidate, 2)
	if o.Count() != 0 {
		t.Fatal("expect Count() to be zero")
	}
	o.queue <- Candidate{}
	if o.Count() != 1 {
		t.Fatalf("expect Count() to be 01(push not used). got: %d", o.Count())
	}
	_ = <-o.queue
	if o.Count() != 0 {
		t.Fatal("expect Count() to be zero again")
	}
}
