package environment

import (
	"fmt"
	environment_testing2 "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireBool(t *testing.T) {
	var errMsg error
	environment_testing2.RunGetBoolTest(t, "b0", RequireBool, "true", nil, true)
	environment_testing2.RunGetBoolTest(t, "b1", RequireBool, "false", nil, false)
	environment_testing2.RunGetBoolTest(t, "b2", RequireBool, "TRUE", nil, true)
	environment_testing2.RunGetBoolTest(t, "b3", RequireBool, "FALSE", nil, false)
	environment_testing2.RunGetBoolTest(t, "b4", RequireBool, "T", nil, true)
	environment_testing2.RunGetBoolTest(t, "b5", RequireBool, "F", nil, false)
	environment_testing2.RunGetBoolTest(t, "b6", RequireBool, "t", nil, true)
	environment_testing2.RunGetBoolTest(t, "b7", RequireBool, "f", nil, false)
	environment_testing2.RunGetBoolTest(t, "b8", RequireBool, "1", nil, true)
	environment_testing2.RunGetBoolTest(t, "b9", RequireBool, "0", nil, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing2.RunGetBoolTest(t, "b10", RequireBool, "BadString", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing2.RunGetBoolTest(t, "b11", RequireBool, "1.0", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing2.RunGetBoolTest(t, "b12", RequireBool, "0.0", errMsg, false)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing2.RunRequiredBoolTest(t, "b12", RequireBool, errMsg)
}
