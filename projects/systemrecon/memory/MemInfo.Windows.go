//go:build windows
// +build windows

package systemrecon

/*
 * MemInfo (Windows)
 *
 * MemInfo() for Windows.
 *
 * 	Return a key-value map of memory-related information from a Windows system.
 */
import (
	misc "github.com/sam-caldwell/go/v2/projects/misc/formatting"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os/exec"
)

// MemInfo - Return a key-value map of memory information delimited by a ':'
func MemInfo() (map[string]string, err error) {
	const (
		columnCount = 2
		keyColumn   = 0
		valueColumn = 1
		sortContent = true
	)
	var output []byte

	cmd := exec.Command("systeminfo")

	if output, err = cmd.Output(); err != nil {
		return nil, err
	}

	return misc.TextToKeyValueMap(&output, words.NewLine, words.Colon, sortContent), nil

}
