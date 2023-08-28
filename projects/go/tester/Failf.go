package tester

import (
	"github.com/sam-caldwell/monorepo/projects/go/v2/ansi/color"
)

// Failf - A Fatal error in a test
func (test *T) Failf(format string, msg ...any) *T {
	test.fail++
	defer test.t.FailNow()
	ansi.Red().
		Print("[FAIL]").Space().
		Printf(format, msg...).LF().
		Reset()
	return test
}
