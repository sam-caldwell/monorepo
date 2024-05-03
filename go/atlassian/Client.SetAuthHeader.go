package Atlassian

import (
    "github.com/sam-caldwell/monorepo/go/misc/words"
    "net/http"
)

// SetAuthHeader - Set the Auth header for a given request
func (client *Client) SetAuthHeader(request *http.Request) error {
    //ToDo: Double-check that the apiKey is base64-encoded properly.
    request.Header.Set(words.Authorization, words.Basic+client.GetApiKey())
    return nil
}
