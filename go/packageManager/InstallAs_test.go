package packageManager

import (
	"reflect"
	"testing"
)

func TestInstallAs(t *testing.T) {
	var data InstallAs
	if reflect.TypeOf(data).Size() != 1 {
		t.Fatalf("InstallAs should be 1 byte\n"+
			"Got: %d", reflect.TypeOf(data).Size())
	}

	if Pkg != 0 {
		t.Fatal("Pkg expects 0")
	}
	if Shell != 1 {
		t.Fatal("Shell expects 1")
	}
	if GoGet != 2 {
		t.Fatal("GoGet expects 2")
	}
	if GoInstall != 3 {
		t.Fatal("GoInstall expects 3")
	}
}
