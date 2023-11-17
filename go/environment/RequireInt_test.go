package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireInt(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	environmentTesting.RunGetIntTest(t, "i0", RequireInt, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	environmentTesting.RunGetIntTest(t, "i1", RequireInt, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	environmentTesting.RunGetIntTest(t, "i2", RequireInt, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetIntTest(t, "i3", RequireInt, "BadString", errMsg, 0.0)

	environmentTesting.RunGetIntTest(t, "i4", RequireInt, "-1", nil, -1.0)
	environmentTesting.RunGetIntTest(t, "i5", RequireInt, "0", nil, 0.0)
	environmentTesting.RunGetIntTest(t, "i6", RequireInt, "1", nil, 1.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredIntTest(t, "i7", RequireInt, errMsg)

}
