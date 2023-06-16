package environment

import (
	"fmt"
	"os"
	"testing"
)

func TestSetFloat(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue = 1337.01
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetFloat(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != fmt.Sprintf("%v", testValue) {
		t.Fatalf("expected value was not set")
	}
}
