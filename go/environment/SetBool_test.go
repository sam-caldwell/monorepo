package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestSetBool(t *testing.T) {
	envVar := "TestEnvVar1337"
	var testValue = true
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetBool(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != fmt.Sprintf("%v", testValue) {
		t.Fatalf("expected value was not set")
	}
}
