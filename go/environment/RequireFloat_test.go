package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireFloat(t *testing.T) {
	var errMsg error

	environmentTesting.RunGetFloatTest(t, "f0", RequireFloat, "-3.141529", nil, -3.141529)
	environmentTesting.RunGetFloatTest(t, "f1", RequireFloat, "0.0", nil, 0.0)
	environmentTesting.RunGetFloatTest(t, "f2", RequireFloat, "3.141529", nil, 3.141529)

	errMsg = fmt.Errorf("strconv.ParseFloat: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetFloatTest(t, "f3", RequireFloat, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunGetFloatTest(t, "f4", RequireFloat, "", errMsg, 0.0)
	environmentTesting.RunGetFloatTest(t, "f5", RequireFloat, "-1", nil, -1.0)
	environmentTesting.RunGetFloatTest(t, "f6", RequireFloat, "0", nil, 0.0)
	environmentTesting.RunGetFloatTest(t, "f7", RequireFloat, "1", nil, 1.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredFloatTest(t, "f8", RequireFloat, errMsg)

}
