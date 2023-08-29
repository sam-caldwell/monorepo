package environment

import (
	os2 "github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"testing"
)

func TestSetString(t *testing.T) {
	const envVar = "TestEnvVar1337"
	const testValue = "value01"
	if os2.Getenv(envVar) != "" {
		if err := os2.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetString(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os2.Getenv(envVar); value != testValue {
		t.Fatalf("expected value was not set")
	}
}
