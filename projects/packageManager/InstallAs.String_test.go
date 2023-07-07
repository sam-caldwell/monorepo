package packageManager

import (
	"testing"
)

func TestInstallAsString(t *testing.T) {
	var data InstallAs

	if data.String() != "Package" {
		t.Fatal("package 0")
	}
	if data = Pkg; data.String() != "Package" {
		t.Fatal("package 0")
	}
	if data = Shell; data.String() != "Shell" {
		t.Fatal("Shell 1")
	}
	if data = GoGet; data.String() != "GoGet" {
		t.Fatal("GoGet 2")
	}
	if data = GoInstall; data.String() != "GoInstall" {
		t.Fatal("GoInstall 3")
	}
}
