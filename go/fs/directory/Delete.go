package directory

import (
	"os"
)

type DeleteRule bool

const (
	DeleteIfEmpty   DeleteRule = false
	RecursiveDelete DeleteRule = true
)

// Delete - Remove directory (empty or not) based on a delete rule.
func Delete(path string, deleteContents DeleteRule) (err error) {
	if deleteContents {
		return os.RemoveAll(path)
	} else {
		return os.Remove(path)
	}
}
