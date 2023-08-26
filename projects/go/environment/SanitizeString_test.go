package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestSanitizeString(t *testing.T) {
	var err error
	err = nil // ok
	environment_testing.RunSanitizedStringTest(t, SanitizeString, err, "s0", "test", "^[a-z]{4}$", "test")

	err = fmt.Errorf(errPatternCheckFailed)
	environment_testing.RunSanitizedStringTest(t, SanitizeString, err, "s1", "test1", "^[a-z]{4}$", "test")
	environment_testing.RunSanitizedStringTest(t, SanitizeString, err, "s2", "", "^[a-z]{4}$", "test")
	environment_testing.RunSanitizedStringTest(t, SanitizeString, err, "s2", "t3st", "^[a-z]{4}$", "test")
}
