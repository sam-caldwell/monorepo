package ansi

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/ansi"
)

// Pass - A passing test report
func (test *T) Pass(msg string) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Println(msg).
		Reset()
	return test
}
