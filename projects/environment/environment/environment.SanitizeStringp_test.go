package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeStringp(t *testing.T) {
	var err error
	err = nil // ok
	runSanitizedStringpTest(t, SanitizeStringp, err, "s0", "test", "^[a-z]{4}$", "test")

	err = fmt.Errorf(errPatternCheckFailed)
	runSanitizedStringpTest(t, SanitizeStringp, err, "s1", "test1", "^[a-z]{4}$", "test")
	runSanitizedStringpTest(t, SanitizeStringp, err, "s2", "", "^[a-z]{4}$", "test")
	runSanitizedStringpTest(t, SanitizeStringp, err, "s2", "t3st", "^[a-z]{4}$", "test")
}
