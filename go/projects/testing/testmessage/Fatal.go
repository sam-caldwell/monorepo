package testmessage

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi"
	"os"
)

// Fatal - A Fatal error in a test
func (test *T) Fatal(msg ...any) {
	test.Failf("%v", msg...)

	test.fail++
	defer test.t.FailNow()
	ansi.Red().
		Printf("[FAIL] %v", msg...).LF().
		Reset()
	os.Exit(1)
}
