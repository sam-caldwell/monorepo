package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireBool(t *testing.T) {
	var errMsg error
	environmentTesting.RunGetBoolTest(t, "b0", RequireBool, "true", nil, true)
	environmentTesting.RunGetBoolTest(t, "b1", RequireBool, "false", nil, false)
	environmentTesting.RunGetBoolTest(t, "b2", RequireBool, "TRUE", nil, true)
	environmentTesting.RunGetBoolTest(t, "b3", RequireBool, "FALSE", nil, false)
	environmentTesting.RunGetBoolTest(t, "b4", RequireBool, "T", nil, true)
	environmentTesting.RunGetBoolTest(t, "b5", RequireBool, "F", nil, false)
	environmentTesting.RunGetBoolTest(t, "b6", RequireBool, "t", nil, true)
	environmentTesting.RunGetBoolTest(t, "b7", RequireBool, "f", nil, false)
	environmentTesting.RunGetBoolTest(t, "b8", RequireBool, "1", nil, true)
	environmentTesting.RunGetBoolTest(t, "b9", RequireBool, "0", nil, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetBoolTest(t, "b10", RequireBool, "BadString", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environmentTesting.RunGetBoolTest(t, "b11", RequireBool, "1.0", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environmentTesting.RunGetBoolTest(t, "b12", RequireBool, "0.0", errMsg, false)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredBoolTest(t, "b12", RequireBool, errMsg)
}
