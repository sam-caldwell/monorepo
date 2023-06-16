package environment

import (
	"fmt"
	"testing"
)

func TestGetIntp(t *testing.T) {
	errMsg := fmt.Errorf("error parsing value")
	runIntpTest(t, "i0", GetIntp, "true", errMsg, 0)
	runIntpTest(t, "i1", GetIntp, "false", errMsg, 0)
	runIntpTest(t, "i2", GetIntp, "BadString", errMsg, 0)
	runIntpTest(t, "i3", GetIntp, "", errMsg, 0)

	errMsg = fmt.Errorf(errEnvBoundCheck)
	runIntpTest(t, "i4", GetIntp, "-1", nil, -1)
	runIntpTest(t, "i5", GetIntp, "0", nil, 0)
	runIntpTest(t, "i6", GetIntp, "1",
		fmt.Errorf("value mismatch on testBool6"), 1)
	runIntpTest(t, "i7", GetIntp, "-1.0", nil, 0)
	runIntpTest(t, "i8", GetIntp, "0.0", nil, 0)
	runIntpTest(t, "i9", GetIntp, "1.0", nil, 0)

}
