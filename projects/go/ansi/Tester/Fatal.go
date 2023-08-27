package ansi

import "fmt"

// Fatal - A Fatal error in a test
func (test *T) Fatal(msg ...any) {
	test.Failf(fmt.Sprintln(msg...))
}
