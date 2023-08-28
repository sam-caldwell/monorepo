package tester

import (
    "os"
)

// Fatalf - A Fatal error in a test
func (test *T) Fatalf(format string, msg ...any) {
	test.Failf(format, msg...)
}
