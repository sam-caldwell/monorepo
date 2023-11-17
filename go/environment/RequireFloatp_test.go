package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireFloatp(t *testing.T) {
	environmentTesting.RunGetFloatpTest(t, "f0", RequireFloatp, "-3.141529", nil, -3.141529)
	environmentTesting.RunGetFloatpTest(t, "f1", RequireFloatp, "0.0", nil, 0.0)
	environmentTesting.RunGetFloatpTest(t, "f2", RequireFloatp, "3.141529", nil, 3.141529)

	errMsg := fmt.Errorf("strconv.ParseFloat: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetFloatpTest(t, "f3", RequireFloatp, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunGetFloatpTest(t, "f4", RequireFloatp, "", errMsg, 0.0)
	environmentTesting.RunGetFloatpTest(t, "f5", RequireFloatp, "-1", nil, -1.0)
	environmentTesting.RunGetFloatpTest(t, "f6", RequireFloatp, "0", nil, 0.0)
	environmentTesting.RunGetFloatpTest(t, "f7", RequireFloatp, "1", nil, 1.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredFloatpTest(t, "f8", RequireFloat, errMsg)
}
