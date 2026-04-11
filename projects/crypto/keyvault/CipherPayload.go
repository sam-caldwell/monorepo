package keyvault

// CipherPayload - encrypted in-memory representation of the keyvault contents
//
// When first read into memory, this is the internal representation of the ciphertext.
//
//	(c) 2023 Sam Caldwell.  MIT License
type CipherPayload struct {
	cipher []byte
}
