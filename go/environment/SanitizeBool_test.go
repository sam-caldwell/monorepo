package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestSanitizeBool(t *testing.T) {
	environment_testing.RunGetBoolTest(t, "b0", SanitizeBool, "true", nil, true)
	environment_testing.RunGetBoolTest(t, "b1", SanitizeBool, "false", nil, false)
	environment_testing.RunGetBoolTest(t, "b2", SanitizeBool, "TRUE", nil, true)
	environment_testing.RunGetBoolTest(t, "b3", SanitizeBool, "FALSE", nil, false)
	environment_testing.RunGetBoolTest(t, "b4", SanitizeBool, "T", nil, true)
	environment_testing.RunGetBoolTest(t, "b5", SanitizeBool, "F", nil, false)
	environment_testing.RunGetBoolTest(t, "b6", SanitizeBool, "t", nil, true)
	environment_testing.RunGetBoolTest(t, "b7", SanitizeBool, "f", nil, false)
	environment_testing.RunGetBoolTest(t, "b8", SanitizeBool, "1", nil, true)
	environment_testing.RunGetBoolTest(t, "b9", SanitizeBool, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b10", SanitizeBool, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b11", SanitizeBool, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing.RunGetBoolTest(t, "b12", SanitizeBool, "0.0", errBadInput, false)
}
