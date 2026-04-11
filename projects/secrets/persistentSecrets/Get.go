package persistentSecrets

import (
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

// Get - Return the secret record with a given name.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (vault *Vault) Get(name string) (item *onepassword.Item, err error) {
	client := connect.NewClient(vault.connectHost, vault.token)
	item, err = client.GetItem(name, vault.vault.String())

	return item, err
}
