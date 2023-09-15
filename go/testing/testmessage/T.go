package testmessage

import "testing"

// T - The state of the Ansi Tester
type T struct {
	t    *testing.T
	pass int
	fail int
	skip int
}
