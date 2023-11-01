package hashScanner

/*
 * hashScanner Candidate
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the candidate structure used to queue up
 * a set of candidate byte strings and associated hashes.
 */

type Candidate struct {
	raw  string
	hash string
}
