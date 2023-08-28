package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestRequireString(t *testing.T) {
	environment_testing.RunGetStringTest(t, "s0", RequireString, " ", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environment_testing.RunGetStringTest(t, "s1", RequireString, "0.0", nil, "0.0")
	environment_testing.RunGetStringTest(t, "s2", RequireString, "3.141529", nil, "3.141529")
	environment_testing.RunGetStringTest(t, "s3", RequireString, "random string", nil, "random string")

	environment_testing.RunRequiredStringTest(t, "s4", RequireString, fmt.Errorf(errEnvVarNotFound))
}
