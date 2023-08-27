package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestSanitizeFloat(t *testing.T) {
	var err error
	err = nil // ok
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f0", "-3.141529", err, -3.141529, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f1", "-3.141529", err, -3.141529, 0, 9)

	err = nil // ok
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, 1, 9)
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, -9, -1)

	err = nil // ok
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f3", "3.141529", err, 3.141529, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatTest(t, SanitizeFloat, "f3", "3.141529", err, 3.141529, 0, 3)
}
