package keyvalue

/*
 * KeyValue.FromBytes
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Consume []bytes and parse by lineEnding and columnDelimiter into a keyvalue.Map
 *  internal state.
 *
 *  Note: this is destructive of any existing state
 */

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"strings"
)

// FromBytes - Given a reference to a byte-array, parse by lines and key-value columns, storing internally.
func (kv *KeyValue) FromBytes(data *[]byte, lineEnding, columnDelimiter string) {

	lines := strings.Split(string(*data), lineEnding)

	kv.Initialize(len(lines), overwrite)

	for _, line := range lines {
		fields := strings.SplitN(line, columnDelimiter, columnCount)
		if len(fields) == columnCount {
			kv.data[strings.TrimSpace(fields[keyColumn])] =
				strings.TrimRight(strings.TrimSpace(fields[valueColumn]), words.Period)
		}
	}
}
