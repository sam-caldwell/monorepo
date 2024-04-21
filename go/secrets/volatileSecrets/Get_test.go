package volatileSecrets

import (
	"bytes"
	"testing"
)

func TestGet(t *testing.T) {
	passphrase := []byte("myPassphrase")
	secret := []byte("this is my secret")
	expected := bytes.Clone(secret)
	p := NewPassword(passphrase, secret)
	if actual := p.Get(); actual != string(expected) {
		t.Fatalf("actual secret does not match expected\n"+
			"actual  : %v\n"+
			"expected: %v", actual, string(expected))
	}
}
