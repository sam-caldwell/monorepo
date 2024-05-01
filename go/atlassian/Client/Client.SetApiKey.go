package Client

// SetApiKey - Sanitize and set the api key
func (client *Client) SetApiKey(apiKey string) error {
    return client.apiKey.FromString(&apiKey)
}
