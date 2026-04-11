package Atlassian

// GetApiKey - Return the api key
func (client *Client) GetApiKey() (apiKey string) {

    return client.apiKey.Get()

}
