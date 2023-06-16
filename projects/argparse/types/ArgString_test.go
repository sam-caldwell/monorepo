package types

import "testing"

func TestArgString(t *testing.T) {
	var s ArgString
	if s != "" {
		t.Fatal("expect initial empty string")
	}
	s = "test"
	if s != "test" {
		t.Fatal("expected 'test' string state")
	}
}
