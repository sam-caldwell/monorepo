package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireString(t *testing.T) {
	environmentTesting.RunGetStringTest(t, "s0", RequireString, " ", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	environmentTesting.RunGetStringTest(t, "s1", RequireString, "0.0", nil, "0.0")
	environmentTesting.RunGetStringTest(t, "s2", RequireString, "3.141529", nil, "3.141529")
	environmentTesting.RunGetStringTest(t, "s3", RequireString, "random string", nil, "random string")

	environmentTesting.RunRequiredStringTest(t, "s4", RequireString, fmt.Errorf(errEnvVarNotFound))
}
