package crypto

/*
 * projects/crypto/Output.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides an Output() method which will
 * concatenate our hash values and return the hash
 * result.
 */

import "encoding/binary"

func (hash *Sha256Stream) Output() (out []byte) {
	out = make([]byte, 32)
	for i, h := range hash.h {
		binary.BigEndian.PutUint32(out[i*4:i*4+4], h)
	}
	return out
}
