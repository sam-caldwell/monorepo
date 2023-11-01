package hashScanner

/*
 * hashScanner CandidateQueue.Pop
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

import (
	"os"
)

func (c *CandidateQueue) Pop() *Candidate {
	c.lock.Lock()
	defer c.lock.Unlock()
	item := <-c.queue
	c.count--
	//log.Printf("Pop() count: %d", c.count)
	_ = os.Stdout.Sync()
	return &item
}
