package hashScanner

import "testing"

/*
 * hashScanner Candidate tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines a unit test for th candidate structure
 */

func TestCandidateStruct(t *testing.T) {
	var o Candidate
	if o.raw != "" {
		t.Fatal("initial state (raw) wrong")
	}
	if o.hash != "" {
		t.Fatal("initial state (hash) wrong")
	}
}
