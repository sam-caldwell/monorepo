package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/environment/environment_testing"
	"testing"
)

func TestSanitizeStringp(t *testing.T) {
	var err error
	err = nil // ok
	environment_testing.RunSanitizedStringpTest(t, SanitizeStringp, err, "s0", "test", "^[a-z]{4}$", "test")

	err = fmt.Errorf(errPatternCheckFailed)
	environment_testing.RunSanitizedStringpTest(t, SanitizeStringp, err, "s1", "test1", "^[a-z]{4}$", "test")
	environment_testing.RunSanitizedStringpTest(t, SanitizeStringp, err, "s2", "", "^[a-z]{4}$", "test")
	environment_testing.RunSanitizedStringpTest(t, SanitizeStringp, err, "s2", "t3st", "^[a-z]{4}$", "test")
}
