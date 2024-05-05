package Atlassian

// GetDebug - Return the client's internal debug state
func (client *Client) GetDebug() bool {
	return client.debug
}
