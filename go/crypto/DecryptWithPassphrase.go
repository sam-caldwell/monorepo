package crypto

import (
	helper "github.com/ProtonMail/gopenpgp/v2/helper"
	"runtime"
)

// DecryptWithPassphrase - Decrypt a cipher text with a given passphrase.
func DecryptWithPassphrase(cipher, passphrase string) (clearText string, err error) {
	// just an abstraction point from upstream
	defer runtime.GC()
	return helper.DecryptMessageWithPassword([]byte(passphrase), cipher)
}
