package testmessage

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Skip - A skipped test report
func (test *T) Skip(msg string) *T {
	test.skip++
	defer test.t.SkipNow()
	ansi.Yellow().
		Print("[SKIP]").Space().
		Println(msg).
		Reset()
	return test
}
