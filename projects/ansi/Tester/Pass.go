package ansi

import "github.com/sam-caldwell/go/v2/projects/ansi"

// Pass - A passing test report
func (test *T) Pass(msg string) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Println(msg).
		Reset()
	return test
}
