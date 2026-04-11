package persistentSecrets

// NewConnection - Create and configure the Vault and its connection to 1Password.
//
//	 This function creates a secure connection to 1Password, where the connection
//	 information is stored in a PGP encrypted file on the local computer.  This
//	 connection information is stored in memory as ciphertext and passed to a
//	 caller upon request
//
//			(c) 2023 Sam Caldwell.  MIT License
func NewConnection() *Vault {
	var vault Vault
	vault.LoadData()
	return &vault
}

func (vault *Vault) connectHost() string {
	return vault.decrypt(0)
}
func (vault *Vault) token() string {
	return vault.decrypt(1)
}
func (vault *Vault) connectToken() string {
	return vault.decrypt(2)
}
