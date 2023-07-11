package simpleArgs

import (
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"strings"
)

// HasFlag - For a given name return a boolean if this name exists in os.Args[2:]
func HasFlag(name string) bool {
	if len(os.Args) < 3 {
		return false
	}
	for _, option := range os.Args[2:] {
		if option == strings.TrimSpace(strings.ToLower(name)) {
			return true
		}
	}
	return false
}
