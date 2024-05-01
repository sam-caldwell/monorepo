package Atlassian

// Init - Initialize the client with its basic parameters
func (client *Client) Init(domain, apiKey string) *Client {

    if client.err = client.domain.Set(&domain); client.err != nil {
        return client
    }
    client.err = client.apiKey.FromString(&apiKey)
    return client

}
