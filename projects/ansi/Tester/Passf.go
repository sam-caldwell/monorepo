package ansi

import "github.com/sam-caldwell/go/v2/projects/ansi"

// Passf - A passing test report
func (test *T) Passf(format string, msg ...any) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Printf(format, msg...).
		Reset()
	return test
}
