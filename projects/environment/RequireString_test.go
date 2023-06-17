package environment

import (
	"fmt"
	"testing"
)

func TestRequireString(t *testing.T) {
	runStringTest(t, "s0", RequireString, "", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	runStringTest(t, "s1", RequireString, " ", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	runStringTest(t, "s2", RequireString, "0.0", nil, "0.0")
	runStringTest(t, "s3", RequireString, "3.141529", nil, "3.141529")
	runStringTest(t, "s4", RequireString, "random string", nil, "random string")
}
