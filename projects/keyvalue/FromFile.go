package keyvalue

/*
 * keyvalue.FromFile()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Read a key-value file and store the contents in the keyvalue Struct
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"strings"
)

// FromFile - Read a key-value file and process it into the KeyValue struct.
func (kv *KeyValue) FromFile(fileName, columnDelimiter, lineEnding string) (err error) {
	var data []byte

	if !file.Exists(fileName) {
		return fmt.Errorf("file not found (%s)", fileName)
	}

	data, err = os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	lines := strings.Split(string(data), lineEnding)

	kv.Initialize(len(lines), overwrite)

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue //ignore commented lines
		}
		fields := strings.SplitN(strings.TrimSpace(line), columnDelimiter, columnCount)

		//Note: we only keep things with 2 columns
		if len(fields) == columnCount {
			lhs := strings.TrimSpace(fields[keyColumn])
			rhs := strings.Split(strings.TrimSpace(fields[valueColumn]), " #")
			kv.data[lhs] = rhs[0]
		}
	}
	return nil
}
