package persistentSecrets

import (
	"fmt"
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// GetSecret - Return a specific secret value given its section and field name.
//
//	 This is a very opinionated variant of Get()
//
//		(c) 2023 Sam Caldwell.  MIT License
func (vault *Vault) GetSecret(name, section, field string) (value string, err error) {
	var item *onepassword.Item
	client := connect.NewClient(vault.connectHost, vault.token)
	item, err = client.GetItem(name, vault.vault.String())
	if err != nil {
		return words.EmptyString, err
	}
	for _, currentField := range item.Fields {
		if currentField.Section.Label == section && currentField.Label == field {
			return currentField.Value, nil
		}
	}
	return value, fmt.Errorf(errors.NotFound)
}
