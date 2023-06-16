package environment

import (
	"fmt"
	"testing"
)

func TestGetFloat(t *testing.T) {
	runFloatTest(t, "f0", GetFloat, "-3.141529", nil, -3.141529)
	runFloatTest(t, "f1", GetFloat, "0.0", nil, 0.0)
	runFloatTest(t, "f2", GetFloat, "3.141529", nil, 3.141529)

	expectedErr := fmt.Errorf("strconv.ParseFloat: parsing \"fuckedUpBadString\": invalid syntax")
	runFloatTest(t, "f3", GetFloat, "fuckedUpBadString", expectedErr, 0.0)

	expectedErr = fmt.Errorf("strconv.ParseFloat: parsing \"\": invalid syntax")
	runFloatTest(t, "f4", GetFloat, "", expectedErr, 0.0)

	runFloatTest(t, "f5", GetFloat, "-1", nil, -1.0)
	runFloatTest(t, "f6", GetFloat, "0", nil, 0.0)
	runFloatTest(t, "f7", GetFloat, "1", nil, 1.0)
}
