package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireStringp(t *testing.T) {
	environmentTesting.RunGetStringpTest(t, "s0", RequireStringp, "", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environmentTesting.RunGetStringpTest(t, "s1", RequireStringp, "0.0", nil, "0.0")
	environmentTesting.RunGetStringpTest(t, "s2", RequireStringp, "3.141529", nil, "3.141529")
	environmentTesting.RunGetStringpTest(t, "s3", RequireStringp, "random string", nil, "random string")

	environmentTesting.RunRequiredStringpTest(t, "s4", RequireStringp, fmt.Errorf(errEnvVarNotFound))
}
