package hashScanner

/*
 * hashScanner CandidateQueue.Count
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

func (c *CandidateQueue) Count() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.queue == nil {
		return 0
	}
	return int64(len(c.queue))
}
