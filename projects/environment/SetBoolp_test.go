package environment

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"testing"
)

func TestSetBoolp(t *testing.T) {
	envVar := "TestEnvVar1337"
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetBoolp(&envVar, true); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != fmt.Sprintf("%v", true) {
		t.Fatalf("expected value was not set")
	}
}
