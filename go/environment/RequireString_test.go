package environment

import (
	"fmt"
	environment_testing2 "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireString(t *testing.T) {
	environment_testing2.RunGetStringTest(t, "s0", RequireString, " ", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environment_testing2.RunGetStringTest(t, "s1", RequireString, "0.0", nil, "0.0")
	environment_testing2.RunGetStringTest(t, "s2", RequireString, "3.141529", nil, "3.141529")
	environment_testing2.RunGetStringTest(t, "s3", RequireString, "random string", nil, "random string")

	environment_testing2.RunRequiredStringTest(t, "s4", RequireString, fmt.Errorf(errEnvVarNotFound))
}
