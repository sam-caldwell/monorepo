package Atlassian

import (
	"fmt"
	env "github.com/sam-caldwell/monorepo/go/environment"
)

// SetApiKey - Sanitize and set the api key
func (client *Client) SetApiKey(apiKey *string) (err error) {

	if (apiKey == nil) || (*apiKey == "") {
		return fmt.Errorf("empty or nil apiKey")
	}

	*apiKey, err = env.RequireString("ATLASSIAN_TOKEN")

	if err != nil {

		return err

	}

	return client.apiKey.Set(apiKey)

}
