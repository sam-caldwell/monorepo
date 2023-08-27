package repoScanner

import (
	"github.com/sam-caldwell/go/v2/projects/go/simpleLogger"
)

// Scan - Run all security scans for the repo or designated project
func Scan(logf simpleLogger.Logf, noop bool) error {
	if noop {
		return nil
	}
	return nil
}
