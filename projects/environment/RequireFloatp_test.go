package environment

import (
	"fmt"
	"testing"
)

func TestRequireFloatp(t *testing.T) {
	runFloatpTest(t, "f0", RequireFloatp, "-3.141529", nil, -3.141529)
	runFloatpTest(t, "f1", RequireFloatp, "0.0", nil, 0.0)
	runFloatpTest(t, "f2", RequireFloatp, "3.141529", nil, 3.141529)

	errMsg := fmt.Errorf("strconv.ParseFloat: parsing \"BadString\": invalid syntax")
	runFloatpTest(t, "f3", RequireFloatp, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	runFloatpTest(t, "f4", RequireFloatp, "", errMsg, 0.0)
	runFloatpTest(t, "f5", RequireFloatp, "-1", nil, -1.0)
	runFloatpTest(t, "f6", RequireFloatp, "0", nil, 0.0)
	runFloatpTest(t, "f7", RequireFloatp, "1", nil, 1.0)
}
