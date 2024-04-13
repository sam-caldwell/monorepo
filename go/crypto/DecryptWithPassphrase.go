package crypto

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"runtime"
)

// DecryptWithPassphrase - Decrypt a cipher text with a given passphrase.
func DecryptWithPassphrase(cipher, passphrase []byte) ([]byte, error) {
	defer runtime.GC()
	clearText, err := helper.DecryptMessageWithPassword(passphrase, string(cipher))
	return []byte(clearText), err
}
