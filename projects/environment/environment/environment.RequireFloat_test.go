package environment

import (
	"fmt"
	"testing"
)

func TestRequireFloat(t *testing.T) {
	runFloatTest(t, "f0", RequireFloat, "-3.141529", nil, -3.141529)
	runFloatTest(t, "f1", RequireFloat, "0.0", nil, 0.0)
	runFloatTest(t, "f2", RequireFloat, "3.141529", nil, 3.141529)

	errMsg := fmt.Errorf("strconv.ParseFloat: parsing \"BadString\": invalid syntax")
	runFloatTest(t, "f3", RequireFloat, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	runFloatTest(t, "f4", RequireFloat, "", errMsg, 0.0)
	runFloatTest(t, "f5", RequireFloat, "-1", nil, -1.0)
	runFloatTest(t, "f6", RequireFloat, "0", nil, 0.0)
	runFloatTest(t, "f7", RequireFloat, "1", nil, 1.0)
}
