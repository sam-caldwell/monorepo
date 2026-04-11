package Atlassian

import (
	"fmt"
	"net/http"
)

// SetAuthHeader - Set the Auth header for a given request
func (client *Client) SetAuthHeader(request *http.Request) error {

	if client.username == "" {
		return fmt.Errorf("empty username")
	}
	if client.GetApiKey() == "" {
		return fmt.Errorf("empty apiKey")
	}

	request.SetBasicAuth(client.username, client.GetApiKey())

	return nil
}
