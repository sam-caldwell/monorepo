package hashScanner

import "sync"

/*
 * hashScanner CandidateQueue
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

type CandidateQueue struct {
	lock  sync.Mutex
	queue chan Candidate
}
