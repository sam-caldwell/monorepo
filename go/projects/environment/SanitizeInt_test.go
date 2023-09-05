package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/environment/environment_testing"
	"testing"
)

func TestSanitizeInt(t *testing.T) {
	var err error
	err = nil // ok
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f0", "-3", err, -3, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f1", "-3", err, -3, 0, 9)

	err = nil // ok
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f2", "0", err, 0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f2", "0", err, 0, 1, 9)
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f2", "0", err, 0, -9, -1)

	err = nil // ok
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f3", "3", err, 3, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedIntTest(t, SanitizeInt, "f3", "3", err, 3, 0, 3)
}
