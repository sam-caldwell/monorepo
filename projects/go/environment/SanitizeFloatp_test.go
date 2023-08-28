package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestSanitizeFloatp(t *testing.T) {
	var err error
	err = nil // ok
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f0", "-3.141529", err, -3.141529, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f1", "-3.141529", err, -3.141529, 0, 9)

	err = nil // ok
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, 1, 9)
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, -9, -1)

	err = nil // ok
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f3", "3.141529", err, 3.141529, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	environment_testing.RunSanitizedFloatpTest(t, SanitizeFloatp, "f3", "3.141529", err, 3.141529, 0, 3)
}
