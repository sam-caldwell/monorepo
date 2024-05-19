package keyvalue

import (
	"strings"
)

// ToString - Return a flattened set (string) using the given rowFormat
//
//	     Given an internal key-value data set, this method will return a string
//	     table of key-value data where each row is formatted using a column delimiter
//	     and line separator.
//
//		 (c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) ToString(columnDelimiter, lineEnding string) (output string) {
	const pretty = true
	return strings.Join(kv.ToStringArray(columnDelimiter, lineEnding, pretty), "")

}
