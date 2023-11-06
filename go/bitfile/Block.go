package bitfile

/*
 * Block
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * The Block allows a sequence of related bytes to be
 * read from a file into a single struct for manipulation.
 */

type Block struct {
	buffer []byte
}
