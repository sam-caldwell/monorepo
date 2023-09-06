package environment

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestSetAny(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue any = "value01"
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetAny(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != testValue {
		t.Fatalf("expected value was not set")
	}
}
