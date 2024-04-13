package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptWithPassphrase(t *testing.T) {
	for i := 0; i < 10; i++ {
		passphrase, _ := GenerateRandomBytes(16)
		plainText, _ := GenerateRandomBytes(4096)
		cipherText, err := EncryptWithPassphrase(plainText, passphrase)
		if err != nil {
			t.Fatal(err)
		}
		result, err := DecryptWithPassphrase(cipherText, passphrase)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Equal(result, plainText) {
			t.Fatalf("Mismatch.\nexpected: '%02x'\nactual: '%02x'", plainText, result)
		}
	}
}
