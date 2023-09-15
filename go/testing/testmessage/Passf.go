package testmessage

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Passf - A passing test report
func (test *T) Passf(format string, msg ...any) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Printf(format, msg...).
		Reset()
	return test
}
