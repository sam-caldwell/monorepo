package environment

import (
	"github.com/sam-caldwell/go/v2/projects/go/environment/environment_testing"
	"testing"
)

func TestGetStringp(t *testing.T) {
	environment_testing.RunGetStringpTest(t, "TestBool0", GetStringp, "true", nil, "true")
	environment_testing.RunGetStringpTest(t, "TestBool1", GetStringp, "false", nil, "false")
	environment_testing.RunGetStringpTest(t, "TestBool2", GetStringp, "some string", nil, "some string")
	environment_testing.RunGetStringpTest(t, "TestBool3", GetStringp, "", nil, "")
}
