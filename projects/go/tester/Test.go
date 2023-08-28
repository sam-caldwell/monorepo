package ansi

import (
	"github.com/sam-caldwell/monorepo/projects/go/v2/ansi/color"
	"testing"
)

/*
 * Usage:
 *   defer ansi.Test(t).Stats()
 *   ansi.Test(t).Pass("this works")
 *   ansi.test(t).Fail("this was broken")
 */

// Test - create a new ansi Tester (T) object
func Test(t *testing.T) *T {
	test := T{t: t}
	ansi.Reset()
	return &test
}
