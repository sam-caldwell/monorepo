package environment

import (
	os2 "github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
	"testing"
)

func TestSetStringp(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue = "value01"
	if os2.Getenv(envVar) != "" {
		if err := os2.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetStringp(&envVar, &testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os2.Getenv(envVar); value != testValue {
		t.Fatalf("expected value was not set")
	}
}