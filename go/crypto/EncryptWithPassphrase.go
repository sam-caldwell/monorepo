package crypto

import (
	helper "github.com/ProtonMail/gopenpgp/v2/helper"
)

// EncryptWithPassphrase - Encrypt a clearText string with a given passphrase.
func EncryptWithPassphrase(clearText, passphrase string) (cipher string, err error) {
	// just an abstraction point from upstream
	return helper.EncryptMessageWithPassword([]byte(passphrase), clearText)
}
