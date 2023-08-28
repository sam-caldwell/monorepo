package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestGetBool(t *testing.T) {
	environment_testing.RunGetBoolTest(t, "b0", GetBool, "true", nil, true)
	environment_testing.RunGetBoolTest(t, "b1", GetBool, "false", nil, false)
	environment_testing.RunGetBoolTest(t, "b2", GetBool, "TRUE", nil, true)
	environment_testing.RunGetBoolTest(t, "b3", GetBool, "FALSE", nil, false)
	environment_testing.RunGetBoolTest(t, "b4", GetBool, "T", nil, true)
	environment_testing.RunGetBoolTest(t, "b5", GetBool, "F", nil, false)
	environment_testing.RunGetBoolTest(t, "b6", GetBool, "t", nil, true)
	environment_testing.RunGetBoolTest(t, "b7", GetBool, "f", nil, false)
	environment_testing.RunGetBoolTest(t, "b8", GetBool, "1", nil, true)
	environment_testing.RunGetBoolTest(t, "b9", GetBool, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b10", GetBool, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b11", GetBool, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b12", GetBool, "0.0", errBadInput, false)
}
