package environment

import (
	"fmt"
	environment_testing2 "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireStringp(t *testing.T) {
	environment_testing2.RunGetStringpTest(t, "s0", RequireStringp, "", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environment_testing2.RunGetStringpTest(t, "s1", RequireStringp, "0.0", nil, "0.0")
	environment_testing2.RunGetStringpTest(t, "s2", RequireStringp, "3.141529", nil, "3.141529")
	environment_testing2.RunGetStringpTest(t, "s3", RequireStringp, "random string", nil, "random string")

	environment_testing2.RunRequiredStringpTest(t, "s4", RequireStringp, fmt.Errorf(errEnvVarNotFound))
}
