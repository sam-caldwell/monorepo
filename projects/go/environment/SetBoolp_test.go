package environment

import (
	"fmt"
	os2 "github.com/sam-caldwell/go/v2/projects/go/wrappers/os"
	"testing"
)

func TestSetBoolp(t *testing.T) {
	envVar := "TestEnvVar1337"
	if os2.Getenv(envVar) != "" {
		if err := os2.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetBoolp(&envVar, true); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os2.Getenv(envVar); value != fmt.Sprintf("%v", true) {
		t.Fatalf("expected value was not set")
	}
}
