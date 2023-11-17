package environment

import (
	"fmt"
	environmentTesting "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireIntp(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	environmentTesting.RunGetIntpTest(t, "i0", RequireIntp, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	environmentTesting.RunGetIntpTest(t, "i1", RequireIntp, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	environmentTesting.RunGetIntpTest(t, "i2", RequireIntp, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	environmentTesting.RunGetIntpTest(t, "i3", RequireIntp, "BadString", errMsg, 0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunGetIntpTest(t, "i4", RequireIntp, "", errMsg, 0)

	environmentTesting.RunGetIntpTest(t, "i5", RequireIntp, "-1", nil, -1)
	environmentTesting.RunGetIntpTest(t, "i6", RequireIntp, "0", nil, 0)
	environmentTesting.RunGetIntpTest(t, "i7", RequireIntp, "1", nil, 1)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environmentTesting.RunRequiredIntpTest(t, "i8", RequireIntp, errMsg)

}
