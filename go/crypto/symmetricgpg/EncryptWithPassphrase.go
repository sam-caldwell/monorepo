package symmetricgpg

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"runtime"
)

// EncryptWithPassphrase - Encrypt a clearText string with a given passphrase.
func EncryptWithPassphrase(clearText, passphrase []byte) ([]byte, error) {
	defer runtime.GC()
	cipher, err := helper.EncryptMessageWithPassword(passphrase, string(clearText))
	return []byte(cipher), err
}
