package environment

import (
	"fmt"
	"testing"
)

func TestGetIntp(t *testing.T) {
	var errMsg error

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"true\": invalid syntax)")
	runIntpTest(t, "i0", GetIntp, "true", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"false\": invalid syntax)")
	runIntpTest(t, "i1", GetIntp, "false", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"BadString\": invalid syntax)")
	runIntpTest(t, "i2", GetIntp, "BadString", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"\": invalid syntax)")
	runIntpTest(t, "i3", GetIntp, "", errMsg, 0)

	runIntpTest(t, "i4", GetIntp, "-1", nil, -1)
	runIntpTest(t, "i5", GetIntp, "0", nil, 0)
	runIntpTest(t, "i6", GetIntp, "1", nil, 1)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"-1.0\": invalid syntax)")
	runIntpTest(t, "i7", GetIntp, "-1.0", errMsg, 0)
	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"0.0\": invalid syntax)")
	runIntpTest(t, "i8", GetIntp, "0.0", errMsg, 0)
	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"1.0\": invalid syntax)")
	runIntpTest(t, "i9", GetIntp, "1.0", errMsg, 0)

}
