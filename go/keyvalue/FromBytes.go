package keyvalue

/*
 * keyvalue.FromBytes
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Consume []bytes and parse by lineEnding and columnDelimiter into a keyvalue.Map
 *  internal state.
 *
 *  Note: this is destructive of any existing state
 */

import (
	"fmt"
	"strings"
)

// FromBytes - Given a reference to a byte-array, parse by lines and key-value columns, storing internally.
//
//	data *[]byte represents the raw data, which we will turn into a string and parse.
//	The lineEnding is any string representing a line delimiter.
//	The columnDelimiter is any string representing a column delimiter.
func (kv *KeyValue[KeyType, ValueType]) FromBytes(data *[]byte, lineEnding, columnDelimiter string) error {
	lines := strings.Split(string(*data), lineEnding)

	kv.Initialize(len(lines), overwrite)

	for lineNumber, line := range lines {
		fields := strings.SplitN(strings.TrimSpace(line), columnDelimiter, columnCount)

		//Note: we only keep things with 2 columns (key and value)
		if len(fields) == columnCount {
			var key KeyType
			var value ValueType

			keyStr := fields[keyColumn]
			valueStr := strings.Join(fields[valueColumn:], columnDelimiter)

			// Use fmt.Sscanf to convert the string to the appropriate type
			if _, err := fmt.Sscanf(keyStr, "%v", &key); err != nil {
				return fmt.Errorf("Error parsing key(%d): %v\n", lineNumber, err)
			}

			if _, err := fmt.Sscanf(valueStr, "%v", &value); err != nil {
				return fmt.Errorf("Error parsing key(%d): %v\n", lineNumber, err)
			}

			kv.lock.Lock()
			kv.data[key] = value
			kv.lock.Unlock()
		}
	}
	return nil
}
