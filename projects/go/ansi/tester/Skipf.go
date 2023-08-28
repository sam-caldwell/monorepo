package ansi

import (
	"github.com/sam-caldwell/monorepo/projects/go/v2/ansi/color"
)

// Skipf - A skipped test report
func (test *T) Skipf(format string, msg ...any) *T {
	test.skip++
	defer test.t.SkipNow()
	ansi.Yellow().
		Print("[SKIP]").Space().
		Printf(format, msg...).
		Reset()
	return test
}
