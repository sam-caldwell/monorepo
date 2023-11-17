package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireBoolp(t *testing.T) {
	var errMsg error

	environmentTesting.RunGetBoolpTest(t, "b0", RequireBoolp, "true", nil, true)
	environmentTesting.RunGetBoolpTest(t, "b1", RequireBoolp, "false", nil, false)
	environmentTesting.RunGetBoolpTest(t, "b2", RequireBoolp, "TRUE", nil, true)
	environmentTesting.RunGetBoolpTest(t, "b3", RequireBoolp, "FALSE", nil, false)
	environmentTesting.RunGetBoolpTest(t, "b4", RequireBoolp, "T", nil, true)
	environmentTesting.RunGetBoolpTest(t, "b5", RequireBoolp, "F", nil, false)
	environmentTesting.RunGetBoolpTest(t, "b6", RequireBoolp, "t", nil, true)
	environmentTesting.RunGetBoolpTest(t, "b7", RequireBoolp, "f", nil, false)
	environmentTesting.RunGetBoolpTest(t, "b8", RequireBoolp, "1", nil, true)
	environmentTesting.RunGetBoolpTest(t, "b9", RequireBoolp, "0", nil, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetBoolpTest(t, "b10", RequireBoolp, "BadString", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environmentTesting.RunGetBoolpTest(t, "b11", RequireBoolp, "1.0", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environmentTesting.RunGetBoolpTest(t, "b12", RequireBoolp, "0.0", errMsg, false)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredBoolTest(t, "b12", RequireBool, errMsg)
}
