package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeString(t *testing.T) {
	var err error
	err = nil // ok
	runSanitizedStringTest(t, SanitizeString, err, "s0", "test", "^[a-z]{4}$", "test")

	err = fmt.Errorf(errPatternCheckFailed)
	runSanitizedStringTest(t, SanitizeString, err, "s1", "test1", "^[a-z]{4}$", "test")
	runSanitizedStringTest(t, SanitizeString, err, "s2", "", "^[a-z]{4}$", "test")
	runSanitizedStringTest(t, SanitizeString, err, "s2", "t3st", "^[a-z]{4}$", "test")
}
