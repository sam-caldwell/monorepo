package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeFloat(t *testing.T) {
	var err error
	err = nil // ok
	runSanitizedFloatTest(t, SanitizeFloat, "f0", "-3.141529", err, -3.141529, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatTest(t, SanitizeFloat, "f1", "-3.141529", err, -3.141529, 0, 9)

	err = nil // ok
	runSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, 1, 9)
	runSanitizedFloatTest(t, SanitizeFloat, "f2", "0.0", err, 0.0, -9, -1)

	err = nil // ok
	runSanitizedFloatTest(t, SanitizeFloat, "f3", "3.141529", err, 3.141529, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedFloatTest(t, SanitizeFloat, "f3", "3.141529", err, 3.141529, 0, 3)
}
