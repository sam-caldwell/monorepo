package volatileSecrets

// Password - stores a password in memory as an encrypted object.
// it's not a perfect secrecy strategy, but it makes it harder to dump secrets from memory
type Password struct {
	data []byte
}
