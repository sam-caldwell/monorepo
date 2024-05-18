package pair

import (
	"bytes"
)

const (
	CompareKeyLess    = -1
	CompareKeyEqual   = 0
	CompareKeyGreater = 1
)

// CompareKey - compare two KeyValue typed values as their byte representations
//
// Possible outcomes:
//
//	 0 if left == right,
//	+1 if left < right, and
//	-1 if left > right
func CompareKey[KeyType comparable](left, right KeyType) int {
	var lhs []byte
	var rhs []byte
	var err error
	getModifier := func(k KeyType) int {
		switch any(k).(type) {
		case int, int8, int16, int32, int64,
			*int, *int8, *int16, *int32, *int64:
			return -1
		default:
			return 1
		}
	}
	if lhs, err = KeyToBytes(left); err != nil {
		panic(err)
	}
	if rhs, err = KeyToBytes(right); err != nil {
		panic(err)
	}
	return getModifier(left) * bytes.Compare(lhs, rhs)
}
