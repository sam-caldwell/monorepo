package crypto

import "sync"

/*
 * projects/crypto/Sha256Stream.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines a struct for the Sha256Stream
 * which allows a user to stream in an arbitrary series
 * of bits or bytes and then compute the hash of the
 * stream.
 */

type Sha256Stream struct {
	// h - An array of hash values a,b,c,d,e,f,g,h in spec.
	h [8]uint32

	// buffer - A 512-bit buffer for our hash blocks.
	//          note that this buffer being defined as a 64-byte
	//          block means we have eliminated the need for an
	//          explicit padding operation since padding is
	//          implied by any unused bytes when the stream
	//          stops.
	buffer [64]byte

	// byteNdx - An index of the current buffer byte element.
	byteNdx int8

	// bitNdx - An index of the current bit in the current byte of the buffer.
	bitNdx int8

	// size - The total bytes processed
	size uint32

	// lock - A lock to make us safer.
	lock sync.Mutex
}
