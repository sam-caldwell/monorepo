package hashScanner

import "testing"

/*
 * hashScanner NewCandidateQueue() function tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the NewCandidateQueue() function tests
 */

func TestNewCandidateQueue(t *testing.T) {
	func() {
		var o CandidateQueue
		if o.queue != nil {
			t.Fatal("expect nil queue initially")
		}
	}()
	func() {
		q := NewCandidateQueue(10)
		if len(q.queue) != 0 {
			t.Fatalf("queue size should be 0. Got: %d", len(q.queue))
		}
		q.queue <- Candidate{}
		if len(q.queue) != 1 {
			t.Fatalf("queue size should be 1. Got: %d", len(q.queue))
		}
		_ = <-q.queue
		if len(q.queue) != 0 {
			t.Fatalf("queue size should be 0. Got: %d", len(q.queue))
		}
	}()
}
