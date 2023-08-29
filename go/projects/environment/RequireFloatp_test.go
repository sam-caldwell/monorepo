package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/environment/environment_testing"
	"testing"
)

func TestRequireFloatp(t *testing.T) {
	environment_testing.RunGetFloatpTest(t, "f0", RequireFloatp, "-3.141529", nil, -3.141529)
	environment_testing.RunGetFloatpTest(t, "f1", RequireFloatp, "0.0", nil, 0.0)
	environment_testing.RunGetFloatpTest(t, "f2", RequireFloatp, "3.141529", nil, 3.141529)

	errMsg := fmt.Errorf("strconv.ParseFloat: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetFloatpTest(t, "f3", RequireFloatp, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing.RunGetFloatpTest(t, "f4", RequireFloatp, "", errMsg, 0.0)
	environment_testing.RunGetFloatpTest(t, "f5", RequireFloatp, "-1", nil, -1.0)
	environment_testing.RunGetFloatpTest(t, "f6", RequireFloatp, "0", nil, 0.0)
	environment_testing.RunGetFloatpTest(t, "f7", RequireFloatp, "1", nil, 1.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing.RunRequiredFloatpTest(t, "f8", RequireFloat, errMsg)
}
