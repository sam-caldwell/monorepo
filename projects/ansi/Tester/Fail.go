package ansi

import "github.com/sam-caldwell/go/v2/projects/ansi"

// Fail - A Fatal error in a test
func (test *T) Fail(msg string) *T {
	test.fail++
	defer test.t.FailNow()
	ansi.Red().
		Print("[FAIL]").Space().
		Println(msg).
		Reset()
	return test
}
