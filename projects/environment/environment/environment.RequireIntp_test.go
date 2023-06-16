package environment

import (
	"fmt"
	"testing"
)

func TestRequireIntp(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	runIntpTest(t, "i0", RequireIntp, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	runIntpTest(t, "i1", RequireIntp, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	runIntpTest(t, "i2", RequireIntp, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	runIntpTest(t, "i3", RequireIntp, "BadString", errMsg, 0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	runIntpTest(t, "i4", RequireIntp, "", errMsg, 0)

	runIntpTest(t, "i5", RequireIntp, "-1", nil, -1)
	runIntpTest(t, "i6", RequireIntp, "0", nil, 0)
	runIntpTest(t, "i7", RequireIntp, "1", nil, 1)
}
