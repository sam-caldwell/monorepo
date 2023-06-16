package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeIntp(t *testing.T) {
	var err error
	err = nil // ok
	runSanitizedIntpTest(t, SanitizeIntp, "f0", "-3", err, -3, -9, 0)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedIntpTest(t, SanitizeIntp, "f1", "-3", err, -3, 0, 9)

	err = nil // ok
	runSanitizedIntpTest(t, SanitizeIntp, "f2", "0", err, 0, -9, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedIntpTest(t, SanitizeIntp, "f2", "0", err, 0, 1, 9)
	runSanitizedIntpTest(t, SanitizeIntp, "f2", "0", err, 0, -9, -1)

	err = nil // ok
	runSanitizedIntpTest(t, SanitizeIntp, "f3", "3", err, 3, 0, 9)

	err = fmt.Errorf(errEnvBoundCheck)
	runSanitizedIntpTest(t, SanitizeIntp, "f3", "3", err, 3, 0, 3)
}
