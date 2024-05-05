package Atlassian

// SetNoop - Set the noop flag
func (client *Client) SetNoop(flag bool) {
	client.noop = flag
}
