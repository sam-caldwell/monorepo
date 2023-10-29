package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter tests
 */

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func TestByteCounter_Sha1(t *testing.T) {
	hashFunc := func(i []byte) string {
		hash := sha1.Sum(i)
		return hex.EncodeToString(hash[:])
	}
	//
	// Happy: Test empty hash
	//
	func() {
		var b ByteCounter
		if b.Sha1() != hashFunc([]byte{}) {
			t.Fatal("sha1 hash of empty state fails")
		}
	}()
	//
	// Happy: Test hash of non-nil set
	//
	func() {
		input := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
		var b ByteCounter
		b.v = input
		if b.Sha1() != hashFunc(input) {
			t.Fatal("mismatch")
		}
	}()
}
