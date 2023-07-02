package ansi

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"strings"
	"testing"
)

/*
 * Usage:
 *   defer ansi.Test(t).Stats()
 *   ansi.Test(t).Pass("this works")
 *   ansi.test(t).Fail("this was broken")
 */

type T struct {
	t    *testing.T
	pass int
	fail int
	skip int
}

func Test(t *testing.T) *T {
	test := T{t: t}
	ansi.Reset()
	return &test
}

func (test *T) Stats() {
	const header = "---Test Statistics---"
	ansi.Bold().Blue().Println(header).Reset().
		Green().Printf("Pass: %d", test.pass).LF().Reset().
		Red().Printf("Fail: %d", test.fail).LF().Reset().
		Yellow().Printf("Skip: %d", test.skip).LF().Reset().
		Bold().Blue().Println(strings.Repeat(words.Hyphen, len(header))).Reset()
}

func (test *T) Pass(msg string) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Println(msg).
		Reset()
	return test
}

func (test *T) Passf(format string, msg ...any) *T {
	test.pass++
	ansi.Green().
		Print("[ OK ]").Space().
		Printf(format, msg...).
		Reset()
	return test
}

func (test *T) Fail(msg string) *T {
	test.fail++
	defer test.t.FailNow()
	ansi.Red().
		Print("[FAIL]").Space().
		Println(msg).
		Reset()
	return test
}

func (test *T) Failf(format string, msg ...any) *T {
	test.fail++
	defer test.t.FailNow()
	ansi.Red().
		Print("[FAIL]").Space().
		Printf(format, msg...).LF().
		Reset()
	return test
}

func (test *T) Fatalf(format string, msg ...any) {
	test.Failf(format, msg...)
}

func (test *T) Fatal(msg ...any) {
	test.Failf(fmt.Sprintln(msg...))
}

func (test *T) Skip(msg string) *T {
	test.skip++
	defer test.t.SkipNow()
	ansi.Yellow().
		Print("[SKIP]").Space().
		Println(msg).
		Reset()
	return test
}

func (test *T) Skipf(format string, msg ...any) *T {
	test.skip++
	defer test.t.SkipNow()
	ansi.Yellow().
		Print("[SKIP]").Space().
		Printf(format, msg...).
		Reset()
	return test
}
