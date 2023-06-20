//go:build linux
// +build linux

package systemrecon

/*
 * MemInfo (Linux)
 *
 * MemInfo() for Linux.
 *
 * 	Return a key-value map of memory-related information from a Linux system.
 */
import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"io/ioutil"
)

// MemInfo - Return a key-value map of memory information delimited by a ':'
func MemInfo() (map[string]string, error) {
	const (
		keyColumn   = 0
		valueColumn = 1
		sortContent = true
		filePath    = "/proc/meminfo"
	)
	var output []byte

	if output, err = ioutil.ReadFile(filePath); err != nil {
		return nil, err
	}

	return misc.TextToKeyValueMap(&output, words.NewLine, words.Colon, sortContent), nil

}
