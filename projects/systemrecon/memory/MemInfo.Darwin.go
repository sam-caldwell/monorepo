//go:build darwin
// +build darwin

package systemrecon

/*
 * MemInfo (MacOS / Darwin)
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * MemInfo() for windows.
 *
 * 	Return a key-value map of memory-related information from a Darwin (macOS) system.
 */
import (
	misc "github.com/sam-caldwell/go/v2/projects/misc/formatting"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os/exec"
)

// MemInfo - Return a key-value map of memory information delimited by a ':'
func MemInfo() (info map[string]string, err error) {
	const (
		keyColumn   = 0
		valueColumn = 1
		sortContent = true
	)

	var output []byte

	cmd := exec.Command("vm_stat")

	if output, err = cmd.Output(); err != nil {
		return nil, err
	}

	return misc.TextToKeyValueMap(&output, words.NewLine, words.Colon, sortContent), nil
}
