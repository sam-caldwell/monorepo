package crypto

/*
 * projects/crypto/String.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides an String() method which will
 * concatenate our hash values and return the hex
 * result as a string.
 */

import "fmt"

func (hash *Sha256Stream) String() (out string) {
	return fmt.Sprintf("%02x", hash.Output())
}
