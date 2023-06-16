package environment

import (
	"fmt"
	"testing"
)

func TestGetFloatp(t *testing.T) {
	runFloatpTest(t, "f0", GetFloatp, "-3.141529", nil, -3.141529)
	runFloatpTest(t, "f1", GetFloatp, "0.0", nil, 0.0)
	runFloatpTest(t, "f2", GetFloatp, "3.141529", nil, 3.141529)
	runFloatpTest(t, "f3", GetFloatp, "fuckedUpBadString",
		fmt.Errorf("strconv.ParseFloat: parsing \"fuckedUpBadString\": invalid syntax"), 0.0)
	runFloatpTest(t, "f4", GetFloatp, "", nil, 0.0)
	runFloatpTest(t, "f5", GetFloatp, "-1", nil, -1.0)
	runFloatpTest(t, "f6", GetFloatp, "0", nil, 0.0)
	runFloatpTest(t, "f7", GetFloatp, "1", nil, 1.0)
}
