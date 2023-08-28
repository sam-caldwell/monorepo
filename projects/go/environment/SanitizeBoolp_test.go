package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestSanitizeBoolp(t *testing.T) {
	environment_testing.RunGetBoolpTest(t, "b0", SanitizeBoolp, "true", nil, true)
	environment_testing.RunGetBoolpTest(t, "b1", SanitizeBoolp, "false", nil, false)
	environment_testing.RunGetBoolpTest(t, "b2", SanitizeBoolp, "TRUE", nil, true)
	environment_testing.RunGetBoolpTest(t, "b3", SanitizeBoolp, "FALSE", nil, false)
	environment_testing.RunGetBoolpTest(t, "b4", SanitizeBoolp, "T", nil, true)
	environment_testing.RunGetBoolpTest(t, "b5", SanitizeBoolp, "F", nil, false)
	environment_testing.RunGetBoolpTest(t, "b6", SanitizeBoolp, "t", nil, true)
	environment_testing.RunGetBoolpTest(t, "b7", SanitizeBoolp, "f", nil, false)
	environment_testing.RunGetBoolpTest(t, "b8", SanitizeBoolp, "1", nil, true)
	environment_testing.RunGetBoolpTest(t, "b9", SanitizeBoolp, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b10", SanitizeBoolp, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b11", SanitizeBoolp, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	environment_testing.RunGetBoolpTest(t, "b12", SanitizeBoolp, "0.0", errBadInput, false)
}
