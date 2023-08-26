package environment

import (
	os2 "github.com/sam-caldwell/go/v2/projects/go/wrappers/os"
	"testing"
)

func TestSetAny(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue any = "value01"
	if os2.Getenv(envVar) != "" {
		if err := os2.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetAny(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os2.Getenv(envVar); value != testValue {
		t.Fatalf("expected value was not set")
	}
}
