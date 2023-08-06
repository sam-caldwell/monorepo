package crypto

import (
	"hash"
)

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
	// The golang hashing library is used because it is optimized for
	// performance.
	h hash.Hash
	// buffer - represents a byte of information.
	buffer byte
	// bitNdx - An index of the current bit in the current byte of the buffer.
	bitNdx int8
}
