package keyvalue

/*
 * keyvalue.FromFile()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Read a key-value file and store the contents in the keyvalue Struct
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"strings"
)

// FromFile - Read a key-value file and process it into the KeyValue struct.
func (kv *KeyValue[KeyType, ValueType]) FromFile(fileName string, colDelimiter, lineEnding, comment rune) (err error) {
	var data []byte

	if !file.Exists(fileName) {
		return fmt.Errorf(errors.NotFound+errors.Details, fileName)
	}

	if data, err = os.ReadFile(fileName); err != nil {
		return fmt.Errorf(errors.CannotReadFile+errors.Details, err)
	}

	lines := strings.Split(string(data), string(lineEnding))

	kv.Initialize(len(lines), overwrite)

	for lineNumber, line := range lines {
		if strings.HasPrefix(line, string(comment)) {
			continue //ignore commented lines
		}
		fields := strings.SplitN(strings.TrimSpace(line), string(colDelimiter), columnCount)

		//Note: we only keep things with 2 columns
		if len(fields) == columnCount {
			var key KeyType
			var value ValueType

			keyStr := fields[keyColumn]
			valueStr := strings.Join(fields[valueColumn:], string(colDelimiter))

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
