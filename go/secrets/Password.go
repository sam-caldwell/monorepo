package secrets

import (
	"github.com/sam-caldwell/monorepo/go/crypto"
	"runtime"
)

// Password - stores a password in memory as an encrypted object.
// it's not a perfect secrecy strategy but it makes it harder to dump secrets from memory
type Password struct {
	data []byte
	key  []byte
}

// NewPassword - create and return a new password object which will keep the secret as a cipher in memory.
// No, this is not perfectly secure.  But it makes it a lot harder to dump the secret.  After all, you don't
// always have to beat the bear, ya just gotta run faster than everyone else running from the bear.
func NewPassword(passphrase []byte, secret []byte) *Password {
	p := Password{}
	p.SetKey(passphrase)
	p.SetBytes(secret)
	return &p
}

// Clear - Overwrite memory and free it.
func (p *Password) Clear() {
	for i := 0; i < len(p.data); i++ {
		p.data[i] = 0x00
	}
	for i := 0; i < len(p.key); i++ {
		p.key[i] = 0x00
	}
	p.data = nil
	p.key = nil
	runtime.GC()
}

// SetKey - Set the encryption key used to obfuscate the secret in memory.
func (p *Password) SetKey(passphrase []byte) {
	p.key = crypto.Sha512Bytes(passphrase)
}

// Set - Set the string in memory
func (p *Password) Set(d string) {
	defer runtime.GC()
	s, err := crypto.EncryptWithPassphrase(d, string(p.key))
	if err != nil {
		panic(err)
	}
	p.data = []byte(s)
}

// SetBytes - Set the byte string in memory.
func (p *Password) SetBytes(d []byte) {
	p.Set(string(d))
}

// Get - Retrieve the secret in memory (yeah...this will put it in cleartext in memory for a time).
// Tip: keep the returned variable local and remove it as soon as possible.
func (p *Password) Get() string {
	s, err := crypto.DecryptWithPassphrase(string(p.data), string(p.key))
	if err != nil {
		panic(err)
	}
	return s
}

// GetBytes - Retrieve the secret in memory (yeah...this will put it in cleartext in memory for a time).
// Tip: keep the returned variable local and remove it as soon as possible.
func (p *Password) GetBytes() []byte {
	s, err := crypto.DecryptWithPassphrase(string(p.data), string(p.key))
	if err != nil {
		panic(err)
	}
	return []byte(s)
}
