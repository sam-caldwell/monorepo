package environment

import (
	"fmt"
	"testing"
)

func TestGetInt(t *testing.T) {
	var errMsg error

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"true\": invalid syntax)")
	runIntTest(t, "i0", GetInt, "true", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"false\": invalid syntax)")
	runIntTest(t, "i1", GetInt, "false", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"BadString\": invalid syntax)")
	runIntTest(t, "i2", GetInt, "BadString", errMsg, 0)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"\": invalid syntax)")
	runIntTest(t, "i3", GetInt, "", errMsg, 0)

	runIntTest(t, "i4", GetInt, "-1", nil, -1)
	runIntTest(t, "i5", GetInt, "0", nil, 0)
	runIntTest(t, "i6", GetInt, "1", nil, 1)

	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"-1.0\": invalid syntax)")
	runIntTest(t, "i7", GetInt, "-1.0", errMsg, 0)
	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"0.0\": invalid syntax)")
	runIntTest(t, "i8", GetInt, "0.0", errMsg, 0)
	errMsg = fmt.Errorf("error parsing value (strconv.ParseInt: parsing \"1.0\": invalid syntax)")
	runIntTest(t, "i9", GetInt, "1.0", errMsg, 0)

}
