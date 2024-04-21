package volatileSecrets

import (
	"bytes"
	"testing"
)

func TestGetBytes(t *testing.T) {
	t.Run("verify that the method returns the decrypted secret", func(t *testing.T) {
		passphrase := []byte("myPassphrase")
		secret := []byte("this is my secret")
		expected := bytes.Clone(secret)
		p := NewPassword(passphrase, secret)
		t.Logf("debugging (test data):\n"+
			"passphrase: %v\n"+
			"  expected: %v", passphrase, expected)

		if actual := p.GetBytes(); string(actual) != string(expected) {
			t.Fatalf("actual secret does not match expected\n"+
				"actual  : %v\n"+
				"expected: %v", actual, expected)
		}
	})

}
