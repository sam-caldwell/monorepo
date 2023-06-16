package environment

import (
	"os"
	"testing"
)

func TestSetStringp(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue = "value01"
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetStringp(&envVar, &testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != testValue {
		t.Fatalf("expected value was not set")
	}
}
