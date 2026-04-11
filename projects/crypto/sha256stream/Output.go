package crypto

/*
 * projects/crypto/Output.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides an Output() method which will
 * concatenate our hash values and return the hash
 * result.
 */

func (hash *Sha256Stream) Output() (out []byte) {
	var input []byte = nil
	if hash.bitNdx != 0 {
		input = []byte{hash.buffer}
	}
	return hash.h.Sum(input)
}
