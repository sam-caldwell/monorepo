package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/environment/environment_testing"
	"testing"
)

func TestRequireStringp(t *testing.T) {
	environment_testing.RunGetStringpTest(t, "s0", RequireStringp, "", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environment_testing.RunGetStringpTest(t, "s1", RequireStringp, "0.0", nil, "0.0")
	environment_testing.RunGetStringpTest(t, "s2", RequireStringp, "3.141529", nil, "3.141529")
	environment_testing.RunGetStringpTest(t, "s3", RequireStringp, "random string", nil, "random string")

	environment_testing.RunRequiredStringpTest(t, "s4", RequireStringp, fmt.Errorf(errEnvVarNotFound))
}
