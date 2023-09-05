package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/environment/environment_testing"
	"testing"
)

func TestGetFloatp(t *testing.T) {
	environment_testing.RunGetFloatpTest(t, "f0", GetFloatp, "-3.141529", nil, -3.141529)
	environment_testing.RunGetFloatpTest(t, "f1", GetFloatp, "0.0", nil, 0.0)
	environment_testing.RunGetFloatpTest(t, "f2", GetFloatp, "3.141529", nil, 3.141529)
	environment_testing.RunGetFloatpTest(t, "f3", GetFloatp, "fuckedUpBadString",
		fmt.Errorf("strconv.ParseFloat: parsing \"fuckedUpBadString\": invalid syntax"), 0.0)
	environment_testing.RunGetFloatpTest(t, "f4", GetFloatp, "", nil, 0.0)
	environment_testing.RunGetFloatpTest(t, "f5", GetFloatp, "-1", nil, -1.0)
	environment_testing.RunGetFloatpTest(t, "f6", GetFloatp, "0", nil, 0.0)
	environment_testing.RunGetFloatpTest(t, "f7", GetFloatp, "1", nil, 1.0)
}
