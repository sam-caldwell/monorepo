package environment

import (
	"fmt"
	environment_testing2 "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireBoolp(t *testing.T) {
	var errMsg error

	environment_testing2.RunGetBoolpTest(t, "b0", RequireBoolp, "true", nil, true)
	environment_testing2.RunGetBoolpTest(t, "b1", RequireBoolp, "false", nil, false)
	environment_testing2.RunGetBoolpTest(t, "b2", RequireBoolp, "TRUE", nil, true)
	environment_testing2.RunGetBoolpTest(t, "b3", RequireBoolp, "FALSE", nil, false)
	environment_testing2.RunGetBoolpTest(t, "b4", RequireBoolp, "T", nil, true)
	environment_testing2.RunGetBoolpTest(t, "b5", RequireBoolp, "F", nil, false)
	environment_testing2.RunGetBoolpTest(t, "b6", RequireBoolp, "t", nil, true)
	environment_testing2.RunGetBoolpTest(t, "b7", RequireBoolp, "f", nil, false)
	environment_testing2.RunGetBoolpTest(t, "b8", RequireBoolp, "1", nil, true)
	environment_testing2.RunGetBoolpTest(t, "b9", RequireBoolp, "0", nil, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing2.RunGetBoolpTest(t, "b10", RequireBoolp, "BadString", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing2.RunGetBoolpTest(t, "b11", RequireBoolp, "1.0", errMsg, false)

	errMsg = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing2.RunGetBoolpTest(t, "b12", RequireBoolp, "0.0", errMsg, false)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing2.RunRequiredBoolTest(t, "b12", RequireBool, errMsg)
}
