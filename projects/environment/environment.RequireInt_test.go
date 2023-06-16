package environment

import (
	"fmt"
	"testing"
)

func TestRequireInt(t *testing.T) {
	errMsg := fmt.Errorf("strconv.ParseInt: parsing \"-3.141529\": invalid syntax")
	runIntTest(t, "i0", RequireInt, "-3.141529", errMsg, -3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"0.0\": invalid syntax")
	runIntTest(t, "i1", RequireInt, "0.0", errMsg, 0)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"3.141529\": invalid syntax")
	runIntTest(t, "i2", RequireInt, "3.141529", errMsg, 3)

	errMsg = fmt.Errorf("strconv.ParseInt: parsing \"BadString\": invalid syntax")
	runIntTest(t, "i3", RequireInt, "BadString", errMsg, 0.0)

	errMsg = fmt.Errorf(errEnvVarNotFound)
	runIntTest(t, "i4", RequireInt, "", errMsg, 0.0)

	runIntTest(t, "i5", RequireInt, "-1", nil, -1.0)
	runIntTest(t, "i6", RequireInt, "0", nil, 0.0)
	runIntTest(t, "i7", RequireInt, "1", nil, 1.0)
}
