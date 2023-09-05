package environment

import (
	"github.com/sam-caldwell/monorepo/go/projects/environment/environment_testing"
	"testing"
)

func TestGetString(t *testing.T) {
	environment_testing.RunGetStringTest(t, "TestBool0", GetString, "true", nil, "true")
	environment_testing.RunGetStringTest(t, "TestBool1", GetString, "false", nil, "false")
	environment_testing.RunGetStringTest(t, "TestBool2", GetString, "some string", nil, "some string")
	environment_testing.RunGetStringTest(t, "TestBool3", GetString, "", nil, "")
}
