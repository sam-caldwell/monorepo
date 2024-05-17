package pair

import "bytes"

// CompareKey - compare two KeyValue typed values as their byte representations
//
// Possible outcomes:
//
//	 0 if left == right,
//	-1 if left < right, and
//	+1 if left > right
func CompareKey[KeyType comparable](left, right KeyType) int {
	var lhs []byte
	var rhs []byte
	var err error
	if lhs, err = KeyToBytes(left); err != nil {
		panic(err)
	}
	if rhs, err = KeyToBytes(right); err != nil {
		panic(err)
	}
	return bytes.Compare(lhs, rhs)
}
