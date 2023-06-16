package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeFloatp(t *testing.T) {
	var err error
	err = nil // ok
	runSanitizedFloatpTest(t, SanitizeFloatp, "f0", "-3.141529", err, -3.141529, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatpTest(t, SanitizeFloatp, "f1", "-3.141529", err, -3.141529, 0, 9)

	err = nil // ok
	runSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, 1, 9)
	runSanitizedFloatpTest(t, SanitizeFloatp, "f2", "0.0", err, 0.0, -9, -1)

	err = nil // ok
	runSanitizedFloatpTest(t, SanitizeFloatp, "f3", "3.141529", err, 3.141529, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatpTest(t, SanitizeFloatp, "f3", "3.141529", err, 3.141529, 0, 3)
}
