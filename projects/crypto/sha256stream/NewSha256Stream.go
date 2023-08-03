package crypto

import "crypto/sha256"

/*
 * projects/crypto/NewSha256Stream
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides NewSha256Stream() which will
 * create and initialize a new sha256 stream object.
 */

// NewSha256Stream - Create and initialize a sha256 stream object.
func NewSha256Stream() *Sha256Stream {
	return &Sha256Stream{
		h:      sha256.New(),
		buffer: 0x00,
		bitNdx: 0,
	}
}
