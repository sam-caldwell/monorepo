package Atlassian

// GetNoop - Return the client's internal noop flag
func (client *Client) GetNoop() bool {
	return client.noop
}
