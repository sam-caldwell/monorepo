package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestRequireInt(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	environment_testing.RunGetIntTest(t, "i0", RequireInt, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	environment_testing.RunGetIntTest(t, "i1", RequireInt, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	environment_testing.RunGetIntTest(t, "i2", RequireInt, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	environment_testing.RunGetIntTest(t, "i3", RequireInt, "BadString", errMsg, 0.0)

	environment_testing.RunGetIntTest(t, "i4", RequireInt, "-1", nil, -1.0)
	environment_testing.RunGetIntTest(t, "i5", RequireInt, "0", nil, 0.0)
	environment_testing.RunGetIntTest(t, "i6", RequireInt, "1", nil, 1.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	environment_testing.RunRequiredIntTest(t, "i7", RequireInt, errMsg)

}
