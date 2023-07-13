package keyvalue

/*
 * keyvalue.ToString
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Given an internal key-value data set, this method will return a string
 * table of key-value data where each row is formatted using a column delimiter
 * and line separator.
 *
 */

import (
	"strings"
)

// ToString - Return a flattened set (string) using the given rowFormat
func (kv *KeyValue) ToString(columnDelimiter string, lineEnding string) (output string) {
	const pretty = true
	return strings.Join(kv.ToStringArray(columnDelimiter, pretty), lineEnding)
}
