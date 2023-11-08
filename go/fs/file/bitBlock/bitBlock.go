package bitBlock

import "sync"

/*
 * bitBlock
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * The Block allows a sequence of related bytes to be
 * read from a file into a single struct for manipulation.
 */

// Block - An atomic block unit.
type Block struct {
	lock   sync.Mutex
	buffer []byte
}
