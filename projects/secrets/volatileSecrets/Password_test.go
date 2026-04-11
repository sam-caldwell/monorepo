package volatileSecrets

import "testing"

func TestPassword_Struct(t *testing.T) {
	var p Password
	if p.data != nil {
		t.Fatal("expect initial state to be nil")
	}
}
