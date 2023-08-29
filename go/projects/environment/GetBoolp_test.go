package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/environment/environment_testing"
	"testing"
)

func TestGetBoolp(t *testing.T) {
	environment_testing.RunGetBoolpTest(t, "b0", GetBoolp, "true", nil, true)
	environment_testing.RunGetBoolpTest(t, "b1", GetBoolp, "false", nil, false)
	environment_testing.RunGetBoolpTest(t, "b2", GetBoolp, "TRUE", nil, true)
	environment_testing.RunGetBoolpTest(t, "b3", GetBoolp, "FALSE", nil, false)
	environment_testing.RunGetBoolpTest(t, "b4", GetBoolp, "T", nil, true)
	environment_testing.RunGetBoolpTest(t, "b5", GetBoolp, "F", nil, false)
	environment_testing.RunGetBoolpTest(t, "b6", GetBoolp, "t", nil, true)
	environment_testing.RunGetBoolpTest(t, "b7", GetBoolp, "f", nil, false)
	environment_testing.RunGetBoolpTest(t, "b8", GetBoolp, "1", nil, true)
	environment_testing.RunGetBoolpTest(t, "b9", GetBoolp, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b10", GetBoolp, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b11", GetBoolp, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b12", GetBoolp, "0.0", errBadInput, false)
}
