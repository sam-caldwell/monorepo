package repoLinter

import "github.com/sam-caldwell/go/v2/projects/simpleLogger"

func Lint(logf simpleLogger.Logf, noop bool) error {
	if noop {
		return nil
	}
	return nil
}
