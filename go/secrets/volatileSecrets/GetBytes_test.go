package volatileSecrets

import (
	"bytes"
	"testing"
)

func TestGetBytes(t *testing.T) {
	passphrase := []byte("myPassphrase")
	secret := []byte("this is my secret")
	p := NewPassword(passphrase, secret)

	t.Run("verify that the method returns the decrypted secret", func(t *testing.T) {
		if actual := p.GetBytes(); !bytes.Equal(actual, secret) {
			t.Fatalf("actual secret does not match expected\n"+
				"actual  : %v\n"+
				"expected: %v", actual, string(secret))
		}
	})

}
