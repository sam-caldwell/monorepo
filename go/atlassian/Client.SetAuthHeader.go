package Atlassian

import (
	"net/http"
)

// SetAuthHeader - Set the Auth header for a given request
func (client *Client) SetAuthHeader(request *http.Request) error {

	request.SetBasicAuth(client.username, client.GetApiKey())

	return nil
}
