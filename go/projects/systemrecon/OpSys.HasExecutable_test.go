package systemrecon

import "testing"

func TestHasExecutable(t *testing.T) {
	expectedCode := 0
	expectedAnswer := "yes"

	actualCode, actualAnswer := HasExecutable("hostname")
	if actualCode != expectedCode {
		t.Fatal("code does not match")
	}
	if expectedAnswer != actualAnswer {
		t.Fatal("answer does not match")
	}
}
