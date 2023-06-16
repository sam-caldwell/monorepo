package types

import "testing"

func TestArgHelp(t *testing.T) {
	var help ArgHelp
	if help != "" {
		t.Fatal("help should be empty")
	}
	help = "test"
	if help != "test" {
		t.Fatal("help should be 'test'")
	}
}
