package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestSetInt(t *testing.T) {
	const envVar = "TestEnvVar1337"
	const testValue = 1337
	if os.Getenv(envVar) != "" {
		if err := os.Unsetenv(envVar); err != nil {
			t.Fail()
		}
	}
	if err := SetInt(envVar, testValue); err != nil {
		t.Fatalf(err.Error())
	}
	if value := os.Getenv(envVar); value != fmt.Sprintf("%v", testValue) {
		t.Fatalf("expected value was not set")
	}
}
