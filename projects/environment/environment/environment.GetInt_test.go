package environment

import (
	"fmt"
	"testing"
)

func TestGetInt(t *testing.T) {
	errMsg := fmt.Errorf("error parsing value")
	runIntTest(t, "i0", GetInt, "true", errMsg, 0)
	runIntTest(t, "i1", GetInt, "false", errMsg, 0)
	runIntTest(t, "i2", GetInt, "BadString", errMsg, 0)
	runIntTest(t, "i3", GetInt, "", errMsg, 0)

	errMsg = fmt.Errorf(errEnvBoundCheck)
	runIntTest(t, "i4", GetInt, "-1", nil, -1)
	runIntTest(t, "i5", GetInt, "0", nil, 0)
	runIntTest(t, "i6", GetInt, "1",
		fmt.Errorf("value mismatch on testBool6"), 1)
	runIntTest(t, "i7", GetInt, "-1.0", nil, 0)
	runIntTest(t, "i8", GetInt, "0.0", nil, 0)
	runIntTest(t, "i9", GetInt, "1.0", nil, 0)

}
