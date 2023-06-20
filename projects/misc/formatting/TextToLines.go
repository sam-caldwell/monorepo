package misc

/*
 * Given a long blob of text consisting of many lines and delimited key-value content
 * such as that in linux /proc/meminfo, parse that information into a map[string]string
 * format and return the map.
 *
 *   - data            - a giant blob of bytes representing string data
 *   - lineEnding      - the character used for lineEndings
 *   - columnDelimiter - the character used to separate key from value
 *   - sortRows        - A boolean (true: sort the lines, false: process as-is)
 */

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"sort"
	"strings"
)

// TextToKeyValueMap - Given a byte array return a parsed key-value map of string data
func TextToKeyValueMap(data *[]byte, lineEnding, columnDelimiter string, sortRows bool) (output map[string]string) {

	const (
		columnCount = 2

		keyColumn = 0

		valueColumn = 1
	)

	lines := strings.Split(string(*data), lineEnding)

	if sortRows {
		sort.Strings(lines)
	}

	output = make(map[string]string, len(lines))

	for _, line := range lines {

		fields := strings.SplitN(line, columnDelimiter, columnCount)

		if len(fields) == columnCount {
			output[strings.TrimSpace(fields[keyColumn])] =
				strings.TrimRight(strings.TrimSpace(fields[valueColumn]), words.Period)
		}

	}

	return output

}
