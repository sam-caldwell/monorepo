package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/environment/environment_testing"
	"testing"
)

func TestGetFloat(t *testing.T) {
	environment_testing.RunGetFloatTest(t, "f0", GetFloat, "-3.141529", nil, -3.141529)
	environment_testing.RunGetFloatTest(t, "f1", GetFloat, "0.0", nil, 0.0)
	environment_testing.RunGetFloatTest(t, "f2", GetFloat, "3.141529", nil, 3.141529)

	expectedErr := fmt.Errorf("strconv.ParseFloat: parsing \"fuckedUpBadString\": invalid syntax")
	environment_testing.RunGetFloatTest(t, "f3", GetFloat, "fuckedUpBadString", expectedErr, 0.0)

	expectedErr = fmt.Errorf("strconv.ParseFloat: parsing \"\": invalid syntax")
	environment_testing.RunGetFloatTest(t, "f4", GetFloat, "", expectedErr, 0.0)

	environment_testing.RunGetFloatTest(t, "f5", GetFloat, "-1", nil, -1.0)
	environment_testing.RunGetFloatTest(t, "f6", GetFloat, "0", nil, 0.0)
	environment_testing.RunGetFloatTest(t, "f7", GetFloat, "1", nil, 1.0)
}
