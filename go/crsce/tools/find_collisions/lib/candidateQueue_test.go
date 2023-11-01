package hashScanner

import "testing"

/*
 * hashScanner CandidateQueue test
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * This file defines the candidateQueue struct test.
 */

func TestCandidateQueueStruct(t *testing.T) {
	var o CandidateQueue

	o.lock.Lock()
	o.lock.Unlock()

	if len(o.queue) != 0 {
		t.Fatal("expect 0-len queue")
	}
}
