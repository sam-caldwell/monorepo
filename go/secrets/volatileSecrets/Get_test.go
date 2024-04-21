package volatileSecrets

import "testing"

func TestGet(t *testing.T) {
	t.Skip("skip while debugging")
	passphrase := []byte("myPassphrase")
	secret := []byte("this is my secret")
	p := NewPassword(passphrase, secret)
	if actual := p.Get(); actual != string(secret) {
		t.Fatalf("actual secret does not match expected\n"+
			"actual  : %v\n"+
			"expected: %v", actual, string(secret))
	}
}
