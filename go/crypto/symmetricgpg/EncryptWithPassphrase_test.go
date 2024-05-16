package symmetricgpg

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/random"
	"testing"
)

func TestEncryptWithPassphrase(t *testing.T) {
	for i := 0; i < 10; i++ {
		passphrase, _ := random.GenerateRandomBytes(16)
		plainText, _ := random.GenerateRandomBytes(4096)
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
