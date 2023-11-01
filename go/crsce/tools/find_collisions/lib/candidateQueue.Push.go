package hashScanner

/*
 * hashScanner CandidateQueue.Push
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

import (
	"os"
)

func (c *CandidateQueue) Push(raw, hash string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.queue == nil {
		//initialize with 10 if not otherwise initialized.
		c.queue = make(chan Candidate, 10)
	}
	c.queue <- Candidate{
		raw:  raw,
		hash: hash,
	}
	//log.Printf("Push() count: %d", c.count)
	_ = os.Stdout.Sync()
}
