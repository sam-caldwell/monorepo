package environment

import (
	"fmt"
	"testing"
)

func TestRequireStringp(t *testing.T) {
	runStringpTest(t, "s0", RequireStringp, "", fmt.Errorf(errEnvVarNotFound), defaultStringValue)
	runStringpTest(t, "s1", RequireStringp, "0.0", nil, "0.0")
	runStringpTest(t, "s2", RequireStringp, "3.141529", nil, "3.141529")
	runStringpTest(t, "s3", RequireStringp, "random string", nil, "random string")
}
