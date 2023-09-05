package environment

import (
	"fmt"
	os2 "github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
	"testing"
)

func TestSetFloat(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue = 1337.01
	if os2.Getenv(envVar) != "" {
		if err := os2.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetFloat(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os2.Getenv(envVar); value != fmt.Sprintf("%v", testValue) {
		t.Fatalf("expected value was not set")
	}
}
