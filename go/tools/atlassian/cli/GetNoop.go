package cli

import (
	"os"
	"strings"
)

// GetNoopFlag - Scan os.Args for --noop and return true.
func GetNoopFlag() bool {
	for arg := range os.Args {
		s := strings.ToLower(os.Args[arg])
		if s == "--noop" {
			return true
		}
	}
	return false
}
