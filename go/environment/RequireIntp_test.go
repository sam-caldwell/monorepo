package environment

import (
	"fmt"
	environment_testing2 "github.com/sam-caldwell/monorepo/go/environment/environment_testing"
	"testing"
)

func TestRequireIntp(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	environment_testing2.RunGetIntpTest(t, "i0", RequireIntp, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	environment_testing2.RunGetIntpTest(t, "i1", RequireIntp, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	environment_testing2.RunGetIntpTest(t, "i2", RequireIntp, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	environment_testing2.RunGetIntpTest(t, "i3", RequireIntp, "BadString", errMsg, 0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing2.RunGetIntpTest(t, "i4", RequireIntp, "", errMsg, 0)

	environment_testing2.RunGetIntpTest(t, "i5", RequireIntp, "-1", nil, -1)
	environment_testing2.RunGetIntpTest(t, "i6", RequireIntp, "0", nil, 0)
	environment_testing2.RunGetIntpTest(t, "i7", RequireIntp, "1", nil, 1)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing2.RunRequiredIntpTest(t, "i8", RequireIntp, errMsg)

}
