package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

// Json - Linter wrapper for JSON files
func Json(filename string) (err error) {
	// This was using jsonlint.  But a year after a PR was sent to Zach Carter with no response
	// This feature is being pulled out.  We will eventually need to fork jsonlint and fix it
	// ourselves (as part of the monorepo?).
	return fmt.Errorf(errors.NotImplemented)
}
