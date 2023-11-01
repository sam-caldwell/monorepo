package hashScanner

/*
 * hashScanner Candidate
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

func NewCandidateQueue(sz int) *CandidateQueue {
	if sz < 0 {
		panic("bounds check error")
	}
	q := CandidateQueue{
		queue: make(chan Candidate, sz),
	}
	return &q
}
