package directory

import (
	"path/filepath"
	"strings"
)

// SplitToList - Given a path (possibly ending in a file name) return a list of only the directories.
func SplitToList(path string) []string {
	thisDir, _ := filepath.Split(path)
	return strings.Split(thisDir, PathSeparator)
}
