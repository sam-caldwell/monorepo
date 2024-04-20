package volatileSecrets

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/crypto"
	"testing"
)

func TestNewPassword(t *testing.T) {
	/*
	 * Starting with these inputs...
	 */
	passphrase := []byte("myPassphrase")
	secret := []byte("this is my secret")
	/*
	 * Calculate our expected encoded forms for passphrase and secret
	 */
	key := crypto.Sha512Bytes(passphrase)
	cipher, _ := crypto.EncryptWithPassphrase(secret, key)
	/*
	 * Create a new secret (encrypted with the passphrase)
	 */
	p := NewPassword(passphrase, secret)
	/*
	 * Verify the secret size uint32 is properly encoded.
	 */
	t.Run("Does the secret size encode properly?", func(t *testing.T) {
		secretSz := binary.BigEndian.Uint32(p.data[0:4])
		expectedSecretSz := uint32(len(cipher))
		if secretSz != expectedSecretSz {
			t.Fatalf("secret size not properly encoded in bytes[0:4]\n"+
				"actualSz  : %d\n"+
				"expectedSz: %d\n"+
				"encoded   : %v"+
				"",
				secretSz, expectedSecretSz, p.data[0:4])
		}
		actualCipher, err := crypto.EncryptWithPassphrase(secret, crypto.Sha512Bytes(passphrase))
		if err != nil {
			panic(err)
		}
		if len(actualCipher) != int(secretSz) {
			t.Fatalf("cipher length is not the calculated size")
		}
	})
	/*
	 *
	 */
	t.Run("confirm that data size matches expectation", func(t *testing.T) {
		//
		// decode the secret size from the bit field
		//
		actualSecretSz := binary.BigEndian.Uint32(p.data[0:4])

		actualSz := len(p.data)

		expectedSz := paddingA + paddingB + paddingC + keySz + uint(actualSecretSz)

		if actualSz != int(expectedSz) {
			keyStart := paddingA - 1
			t.Fatalf("p.data is the wrong size\n"+
				"  Actual  : %d\n"+
				"  Expected: %d\n"+
				"\n"+
				"secretSz   : %d\n"+
				"len(passp) : %d\n"+
				"len(secret): %d\n"+
				"raw:\n"+
				"        sz:%02x\n"+
				"passphrase:%02x\n",
				actualSz, expectedSz, actualSecretSz,
				len(passphrase)*2, len(secret),
				p.data[0:4],
				p.data[keyStart:keyStart+keySz])
		}
	})

}
