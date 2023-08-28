package tester

import (
    "os"
)

// Fatal - A Fatal error in a test
func (test *T) Fatal(msg ...any) {
	test.Failf(msg...)
	os.Exit(1)
}
